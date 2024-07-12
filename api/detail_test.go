package api

import (
	"reflect"
	"testing"
)

func TestGetQuestionDetail(t *testing.T) {
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		args    args
		want    *DataDetail
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{slug: "two-sum"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetQuestionDetail(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuestionDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuestionDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
