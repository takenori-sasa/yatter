package statuses

import (
	"net/http"
	"testing"
	"yatter-backend-go/app/domain/repository"
)

func Test_handler_CreateStatus(t *testing.T) {
	type fields struct {
		sr repository.Status
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				sr: tt.fields.sr,
			}
			h.CreateStatus(tt.args.w, tt.args.r)
		})
	}
}
