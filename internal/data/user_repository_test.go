package data

import (
	"context"
	"reflect"
	"testing"

	"github.com/cmelgarejo/badge-api/pkg/user"
)

func TestUserRepository_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		ur      *UserRepository
		args    args
		want    []user.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetOne(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		ur      *UserRepository
		args    args
		want    user.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetOne(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetByUsername(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		ur      *UserRepository
		args    args
		want    user.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *user.User
	}
	tests := []struct {
		name    string
		ur      *UserRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Create(tt.args.ctx, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
		u   user.User
	}
	tests := []struct {
		name    string
		ur      *UserRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Update(tt.args.ctx, tt.args.id, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		ur      *UserRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
