package result

import (
	"reflect"
	"testing"
)

func TestResponse_ModifyTag(t *testing.T) {
	type args struct {
		tagName string
	}
	want := make(map[string]interface{})
	want["status_code"] = "0"
	want["status_msg"] = "SUCCESS"
	type test struct {
		name    string
		r       *Response
		args    args
		want    map[string]interface{}
		wantErr bool
	}
	tests := []test{
		{
			name: "test empty data",
			r: &Response{
				StatusCode: "0",
				StatusMsg:  "SUCCESS",
				Data:       nil,
			},
			args:    args{"empty data"},
			want:    map[string]interface{}{"status_code": "0", "status_msg": "SUCCESS"},
			wantErr: false,
		},
		{
			name: "test main data",
			r: &Response{
				StatusCode: "0",
				StatusMsg:  "SUCCESS",
				Data:       "main data",
			},
			args:    args{"maindata"},
			want:    map[string]interface{}{"status_code": "0", "status_msg": "SUCCESS", "maindata": "main data"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ModifyTag(tt.args.tagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Response.ModifyTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.ModifyTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTagName(t *testing.T) {
	type no struct {
		Name string
	}
	type exist struct {
		Name string `json:"name"`
	}
	type args struct {
		ptr interface{}
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{
			name: "test no tag struct",
			args: args{&no{}},
			want: "",
		},
		{
			name: "test exist tag struct",
			args: args{&exist{}},
			want: "name",
		},
		{
			name: "test nil pointer",
			args: args{nil},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTagName(tt.args.ptr); got != tt.want {
				t.Errorf("GetTagName() = %v, want %v", got, tt.want)
			}
		})
	}
}
