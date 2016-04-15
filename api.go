package go_todoist

import (
	"net/http"
	"net/url"
	"io"
	"encoding/json"
	"log"
)

const (
	MAX_COMMANDS_PER_REQ = 100
	MAX_NUMBER_OF_REQ_PER_SEC = 50
)

type TodoistAPI struct  {

	ApiEndpoint string
	Token string
	Commands CommandQueue
	TempIdMapping map[string]int `json:"TempIdMapping,omitempty"`
	SyncStatus SyncStatus `json:"SyncStatus,omitempty"`

	DayOrdersTimestamp string `json:"DayOrdersTimestamp,omitempty"`
	UserId int `json:"UserId,omitempty"`
	SeqNo int `json:"seq_no"`
	SeqNoGlobal int `json:"seq_no_global,omitempty"`
	WebStaticVersion int `json:"WebStaticVersion,omitempty"`
	LiveNotificationsLastRead int `json:"LiveNotificationsLastRead,omitempty"`

	User *UserManager `json:"User,omitempty"`
	Projects *ProjectManager `json:"Projects,omitempty"`
	Filters FilterManager `json:"Filters,omitempty"`
	Items ItemManager `json:"Items,omitempty"`
	Reminders ReminderManager `json:"Reminders,omitempty"`
	Labels LabelManager `json:"Labels,omitempty"`
	Notes NoteManager `json:"Notes,omitempty"`
}


type SyncStatus map[string]interface{}

func NewTodoistAPI(token string) *TodoistAPI {
	api := &TodoistAPI{
		Token: token,
		ApiEndpoint: "https://api.todoist.com",
	}

	api.Projects = &ProjectManager{Api:api}

	return api
}

func (api *TodoistAPI) getApiUrl(call string) string {
	return api.ApiEndpoint + "/API/v6/" + call
}

//func (api *TodoistAPI) getSeqNo()

func (api *TodoistAPI) post(call string, data url.Values, v interface{}) error {
	u := api.getApiUrl(call)
	data.Add("token", api.Token)


	resp, err := http.PostForm(u, data)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	if err != nil {
		// return response status code
		return err
	}

	return nil
}

func (api *TodoistAPI) read() error {
	data := url.Values{}

	data.Add("seq_no", "0") // TODO: create seq_no choice
	data.Add("resource_types", "[\"all\"]") //TODO: add resource types choice
	data.Add("day_orders_timestamp", api.DayOrdersTimestamp)


	err := api.post("sync", data, api)

	if err != nil {
		return err
	}

	return nil
}

func (api *TodoistAPI) write() error {
	commands, err := json.Marshal(api.Commands)
	if err != nil {
		return err
	}
	data := url.Values{}
	v := string(commands)
	log.Println(v)
	data.Add("commands", v)

	err = api.post("sync", data, api)

	if err != nil {
		return err
	}
	api.Commands.Clear()

	return nil
}

func (api *TodoistAPI) Sync() error {
	err := api.read()
	if !api.Commands.IsEmpty() {
		err = api.write()
	}
	return err
}
func (api *TodoistAPI) Register(email, fullName, password string) (*User, error) {
	data := url.Values{}

	data.Add("email", email)
	data.Add("full_name", fullName)
	data.Add("password", password)

	user := new(User)
	err := api.post("register", data, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
