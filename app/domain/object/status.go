package object

import (
	"time"
)

// 本当はPasswordHashがハッシュされたパスワードであることを型で保証したい。
// ハッシュ化されたパスワード用の型を用意してstringと区別して管理すると良い。
// 今回は簡単のためstringで管理している。

type Status struct {
	// The internal ID of the account
	ID        int64 `json:"id,omitempty"`
	AccountID int64 `json:"account_id,omitempty"`

	// URL to the avatar image
	Content *string `json:"content,omitempty"`

	// URL to the header image
	URL *string `json:"header,omitempty"`

	// The time the account was created
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}

// // Check if given password is match to account's password
// func (a *Account) CheckPassword(pass string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
// }
//
// // Hash password and set it to account object
// func (a *Account) SetPassword(pass string) error {
// 	passwordHash, err := generatePasswordHash(pass)
// 	if err != nil {
// 		return fmt.Errorf("generate error: %w", err)
// 	}
// 	a.PasswordHash = passwordHash
// 	return nil
// }
//
// func generatePasswordHash(pass string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", fmt.Errorf("hashing password failed: %w", err)
// 	}
// 	return string(hash), nil
// }
