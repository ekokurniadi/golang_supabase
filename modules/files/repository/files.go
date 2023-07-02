package repository

import (
	"encoding/json"
	"golang_drive/core/constant"
	"golang_drive/modules/files/entity"

	supa "github.com/nedpals/supabase-go"
)

type FilesRepository interface {
	GetAllFiles() ([]entity.Files, error)
}

type filesRepository struct {
	supabaseClient *supa.Client
}

func NewFilesRepository(supabaseClient *supa.Client) *filesRepository {
	return &filesRepository{supabaseClient}
}

func (r *filesRepository) GetAllFiles() ([]entity.Files, error) {
	var result []map[string]interface{}
	err := r.supabaseClient.DB.From(constant.TableFiles).Select("*").Execute(&result)

	if err != nil {
		return nil, err
	}

	s := []entity.Files{}

	if len(result) > 0 {
		value, _ := json.Marshal(result)
		_ = json.Unmarshal(value, &s)
	}

	return s, nil
}
