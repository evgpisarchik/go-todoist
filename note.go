package go_todoist

import "encoding/json"

type FileAttachment struct {
	FileName    string `json:"file_name"`
	FileSize    int    `json:"file_size"`
	FileType    string `json:"file_type"`
	FileURL     string `json:"file_url"`
	UploadState string `json:"upload_state"`
}

type Note struct {
	Content      string `json:"content"`
	Attachment   FileAttachment `json:"file_attachment"`
	ID           int         `json:"id"`
	IsArchived   int         `json:"is_archived"`
	IsDeleted    int         `json:"is_deleted"`
	ItemID       int         `json:"item_id"`
	Posted       string      `json:"posted"`
	PostedUID    int         `json:"posted_uid"`
	ProjectID    int         `json:"project_id"`
	UidsToNotify interface{} `json:"uids_to_notify"`
}

func (n Note) String() string {
	return n.Content
}

type NoteManager struct {
	Notes []Note
}

func (nm *NoteManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nm.Notes)
	if err != nil {
		return err
	}

	return nil
}
