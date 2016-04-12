package go_todoist

import (
	"net/http"
	"net/url"
	"io"
	"encoding/json"
)

type TodoistAPI struct  {
	ApiEndpoint string
	Token string
	Commands []Command

	DayOrdersTimestamp string `json:"DayOrdersTimestamp,ommitempty"`
	UserId int `json:"UserId,ommitempty"`
	SeqNo int `json:"seq_no"`
	WebStaticVersion int `json:"WebStaticVersion,ommitempty"`
	LiveNotificationsLastRead int `json:"LiveNotificationsLastRead,ommitempty"`

	User *UserManager `json:"User,ommitempty"`
	Projects *ProjectManager `json:"Projects,ommitempty"`
	Filters FilterManager `json:"Filters,ommitempty"`
	Items ItemManager `json:"Items,ommitempty"`
	Reminders ReminderManager `json:"Reminders,ommitempty"`
	Labels LabelManager `json:"Labels,ommitempty"`
	Notes NoteManager `json:"Notes,ommitempty"`
}

func NewTodoistAPI(token string) *TodoistAPI {
	return &TodoistAPI{
		Token: token,
		ApiEndpoint: "https://api.todoist.com",
	}
}

func (api *TodoistAPI) getApiUrl(call string) string {
	return api.ApiEndpoint + "/API/v6/" + call
}

func (api *TodoistAPI) post(call string, data url.Values, v interface{}) error {
	u := api.getApiUrl(call)

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

func (api *TodoistAPI) Sync() error {
	data := url.Values{}
	//commands, err := json.Marshal(api.Commands)

	//if err != nil {
	//	return err
	//}
	data.Add("token", api.Token)
	//data.Add("commands", string(commands))
	data.Add("seq_no", "0")
	data.Add("resource_types", "[\"all\"]")

	err := api.post("sync", data, api)

	if err != nil {
		return err
	}

	return nil
}
