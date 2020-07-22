package cmd

import (
	"errors"
	"testing"

	"github.com/ZupIT/ritchie-cli/pkg/formula"
	"github.com/ZupIT/ritchie-cli/pkg/github"
	"github.com/ZupIT/ritchie-cli/pkg/prompt"
)

func Test_addRepoCmd_runPrompt(t *testing.T) {
	type fields struct {
		repo               formula.RepositoryAddLister
		github             github.Repositories
		InputTextValidator prompt.InputTextValidator
		InputPassword      prompt.InputPassword
		InputURL           prompt.InputURL
		InputList          prompt.InputList
		InputBool          prompt.InputBool
		InputInt           prompt.InputInt
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Run with success",
			fields: fields{
				repo:               defaultRepoAdderMock,
				github:             defaultGitRepositoryMock,
				InputTextValidator: inputTextValidatorMock{},
				InputPassword:      inputPasswordMock{},
				InputURL:           inputURLMock{},
				InputBool:          inputTrueMock{},
				InputInt:           inputIntMock{},
				InputList:          inputListMock{},
			},
			wantErr: false,
		},
		{
			name: "Fail when repo.List return err",
			fields: fields{
				repo: repoListerAdderCustomMock{
					list: func() (formula.Repos, error) {
						return nil, errors.New("some error")
					},
				},
				github:             defaultGitRepositoryMock,
				InputTextValidator: inputTextValidatorMock{},
				InputPassword:      inputPasswordMock{},
				InputURL:           inputURLMock{},
				InputBool:          inputTrueMock{},
				InputInt:           inputIntMock{},
				InputList:          inputListMock{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewAddRepoCmd(
				tt.fields.repo,
				tt.fields.github,
				tt.fields.InputTextValidator,
				tt.fields.InputPassword,
				tt.fields.InputURL,
				tt.fields.InputList,
				tt.fields.InputBool,
				tt.fields.InputInt,
			)
			o.PersistentFlags().Bool("stdin", false, "input by stdin")
			if err := o.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("init_runPrompt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
