package data

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/cmelgarejo/badge-api/pkg/points"
)

func TestPointRepository_GetPoints(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint
		start  *time.Time
		end    *time.Time
	}
	tests := []struct {
		name    string
		pr      *PointRepository
		args    args
		want    []points.Point
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pr.GetPoints(tt.args.ctx, tt.args.userID, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("PointRepository.GetPoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointRepository.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointRepository_AssignPoints(t *testing.T) {
	type args struct {
		ctx context.Context
		p   *points.Point
	}
	tests := []struct {
		name    string
		pr      *PointRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pr.AssignPoints(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("PointRepository.AssignPoints() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
