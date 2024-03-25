package photo_repository

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
)

type PhotoRepository interface {
	AddPhoto(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error)
	GetPhotos() ([]PhotoUserMapped, errs.Error)
	GetPhotoId(photoId int) (*PhotoUserMapped, errs.Error)
	UpdatePhoto(photoId int, photoPayload *entity.Photo) (*dto.PhotoUpdateResponse, errs.Error)
	DeletePhoto(photoId int) errs.Error
}
