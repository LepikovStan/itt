package usecases

import (
	"reflect"
	"test/domains"
	"test/infrastructure"
	"test/repository"
	"testing"

	"github.com/google/uuid"
)

func TestUseCases_RegisterUser(t *testing.T) {
	successUser := domains.NewUser().SetIsMarried(false).SetLastName("L").SetFirstName("S").SetPassword("123456789").SetAge(21).SetID(uuid.UUID{})

	type fields struct {
		users repository.UserRepository
	}
	type args struct {
		in RegisterUserInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    RegisterUserOutput
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{users: infrastructure.NewUserRepository()},
			args: args{in: RegisterUserInput{
				Age:       21,
				IsMarried: false,
				Firstname: "S",
				Lastname:  "L",
				Fullname:  "S L",
				Password:  "123456789",
			}},
			want: RegisterUserOutput{
				User: successUser,
			},
			wantErr: false,
		},
		{
			name:   "fail user age",
			fields: fields{users: infrastructure.NewUserRepository()},
			args: args{in: RegisterUserInput{
				Age:       16,
				IsMarried: false,
				Firstname: "S",
				Lastname:  "L",
				Fullname:  "S L",
				Password:  "123456789",
			}},
			want:    RegisterUserOutput{},
			wantErr: true,
		},
		{
			name:   "fail user password",
			fields: fields{users: infrastructure.NewUserRepository()},
			args: args{in: RegisterUserInput{
				Age:       21,
				IsMarried: false,
				Firstname: "S",
				Lastname:  "L",
				Fullname:  "S L",
				Password:  "123456",
			}},
			want:    RegisterUserOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UseCases{
				users: tt.fields.users,
			}

			got, err := uc.RegisterUser(tt.args.in)
			got.User = got.User.SetID(uuid.UUID{})
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
