package data

import (
	"context"
	"reflect"
	"testing"

	"github.com/cmelgarejo/badge-api/pkg/badge"
)

func TestBadgeRepository_Create(t *testing.T) {
	type args struct {
		ctx   context.Context
		orgID uint
		b     *badge.Badge
	}
	tests := []struct {
		name    string
		br      *BadgeRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.br.Create(tt.args.ctx, tt.args.orgID, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("BadgeRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBadgeRepository_GetByOrg(t *testing.T) {
	type args struct {
		ctx   context.Context
		orgID uint
	}
	tests := []struct {
		name    string
		br      *BadgeRepository
		args    args
		want    []badge.Badge
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.br.GetByOrg(tt.args.ctx, tt.args.orgID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BadgeRepository.GetByOrg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BadgeRepository.GetByOrg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBadgeRepository_GetByUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint
	}
	tests := []struct {
		name    string
		br      *BadgeRepository
		args    args
		want    []badge.Badge
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.br.GetByUser(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BadgeRepository.GetByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BadgeRepository.GetByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBadgeRepository_AssignBadge(t *testing.T) {
	type args struct {
		ctx     context.Context
		userID  uint
		badgeID uint
	}
	tests := []struct {
		name    string
		br      *BadgeRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.br.AssignBadge(tt.args.ctx, tt.args.userID, tt.args.badgeID); (err != nil) != tt.wantErr {
				t.Errorf("BadgeRepository.AssignBadge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
