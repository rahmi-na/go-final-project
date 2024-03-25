package comment_repository

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
)

type CommentRepository interface {
	AddComment(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error)
	GetComments() ([]CommentUserPhotoMapped, errs.Error)
	GetCommentById(commentId int) (*CommentUserPhotoMapped, errs.Error)
	DeleteComment(commentId int) errs.Error
	UpdateComment(commentId int, commentPayload *entity.Comment) (*entity.Comment, errs.Error)
}
