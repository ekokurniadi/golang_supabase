package service

import (
	"golang_drive/modules/files/entity"
	"golang_drive/modules/files/repository"
)

type FilesService interface {
	GetAllFiles() ([]entity.Files, error)
}

type filesService struct {
	repository repository.FilesRepository
}

func NewFilesService(repository repository.FilesRepository) *filesService {
	return &filesService{repository}
}

func (s *filesService) GetAllFiles() ([]entity.Files, error) {
	result, err := s.repository.GetAllFiles()
	if err != nil {
		return nil, err
	}

	return result, nil
}
