package object

import (
	"time"
)

// Status is the main structure for storing statuses in the system.
// Each status belongs to a particular Account (given by AccountID).
// A Status has an ID, a content, a URL, and a creation timestamp.
// The structure also contains a linked Account object for easier access to related data.
type Status struct {
	// ID is the internal ID of the status.
	ID int64 `json:"id,omitempty"`

	// AccountID is the ID of the account that posted this status.
	AccountID int64 `json:"account_id,omitempty" db:"account_id"`

	// Content is the actual content of the status.
	Content *string `json:"content,omitempty"`

	// URL is a link to the full status (if applicable).
	URL *string `json:"url,omitempty"`

	// CreateAt is the time the status was created.
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`

	// Account is the account that posted this status. It is included for easier access to related data.
	Account *Account `json:"account,omitempty" db:"account"`
}
