package data

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func Test_getConnection(t *testing.T) {
	tests := []struct {
		name    string
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("getConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunMigrations(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunMigrations(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("RunMigrations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
