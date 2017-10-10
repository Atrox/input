package input

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/atrox/homedir"
)

var homeBase string

func init() {
	homeBase, _ = homedir.Expand("~")
}

func TestRequiredValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"empty", "", nil, true},
		{"not empty", "abc", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RequiredValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequiredValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequiredValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"existing dir", "~/.config", filepath.Join(homeBase, ".config"), false},
		{"wrong path", "~/wrong/path", filepath.Join(homeBase, "wrong", "path"), false},
		{"empty without required validator", "", "", false},
		{"not a directory", "~/.gitconfig", filepath.Join(homeBase, ".gitconfig"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDirectoryValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"existing dir", "~/.config", filepath.Join(homeBase, ".config"), false},
		{"wrong path", "~/wrong/path", nil, true},
		{"empty without required validator", "", "", false},
		{"not a directory", "~/.gitconfig", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirectoryValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"existing file", "~/.gitconfig", filepath.Join(homeBase, ".gitconfig"), false},
		{"wrong path", "~/wrong/path", nil, true},
		{"empty without required validator", "", "", false},
		{"not a file", "~/.config", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegerValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"number", "42", 42, false},
		{"not a number", "4n2", nil, true},
		{"empty", "", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntegerValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntegerValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegerValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooleanValidator(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"y", "y", true, false},
		{"yes", "yes", true, false},
		{"n", "n", false, false},
		{"no", "n", false, false},
		{"ignore case sensitivity", "YeS", true, false},
		{"wrong input", "hodor", nil, true},
		{"empty without required validator", "", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BooleanValidator(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("BooleanValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
