package object

// Timeline represents a collection of Status objects.
// It is used to manage and access multiple statuses at once.
// For example, it can represent the timeline of a particular user, consisting of all statuses posted by the user.
type Timeline struct {
	Body []*Status `json:"timeline,omitempty" db:"status"`
}
