package manager

import "be-b-impact.com/csr/usecase"

// UseCaseManager -> all use case
type UseCaseManager interface {
	UsersUseCase() usecase.UsersUseCase
	CategoryUseCase() usecase.CategoryUseCase
	TagUseCase() usecase.TagUseCase
	ContentUseCase() usecase.ContentUseCase
	ImageUseCase() usecase.ImageUseCase
	TagsContentUseCase() usecase.TagsContentUseCase
	ProposalUseCase() usecase.ProposalUseCase
	ProposalDetailUseCase() usecase.ProposalDetailUseCase
	FileUseCase() usecase.FileUseCase
	ProgressUseCase() usecase.ProgressUseCase
	ProposalProgressUseCase() usecase.ProposalProgressUseCase
}

type useCaseManager struct {
	repoManger RepositoryManager
}

// ProgressUseCase implements UseCaseManager.
func (u *useCaseManager) ProgressUseCase() usecase.ProgressUseCase {
	return usecase.NewProgressUseCase(u.repoManger.ProgressRepo())
}

// ProposalProgressUseCase implements UseCaseManager.
func (u *useCaseManager) ProposalProgressUseCase() usecase.ProposalProgressUseCase {
	return usecase.NewProposalProgressUseCase(u.repoManger.ProposalProgressRepo())
}

// FileUseCase implements UseCaseManager.
func (u *useCaseManager) FileUseCase() usecase.FileUseCase {
	return usecase.NewFileUseCase(u.repoManger.FileRepo())
}

// ProposalDetailUseCase implements UseCaseManager.
func (u *useCaseManager) ProposalDetailUseCase() usecase.ProposalDetailUseCase {
	return usecase.NewProposalDetailUseCase(u.repoManger.ProposalDetailRepo())
}

// ProposalUseCase implements UseCaseManager.
func (u *useCaseManager) ProposalUseCase() usecase.ProposalUseCase {
	return usecase.NewProposalUseCase(u.repoManger.ProposalRepo())
}

// TagsContentUseCase implements UseCaseManager.
func (u *useCaseManager) TagsContentUseCase() usecase.TagsContentUseCase {
	return usecase.NewTagsContentUseCase(u.repoManger.TagsContentRepo())
}

// ImageUseCase implements UseCaseManager
func (u *useCaseManager) ImageUseCase() usecase.ImageUseCase {
	return usecase.NewImageUseCase(u.repoManger.ImageRepo())
}

// ContentUseCase implements UseCaseManager
func (u *useCaseManager) ContentUseCase() usecase.ContentUseCase {
	return usecase.NewContentUseCase(u.repoManger.ContentRepo())
}

// TagUseCase implements UseCaseManager
func (u *useCaseManager) TagUseCase() usecase.TagUseCase {
	return usecase.NewTagUseCase(u.repoManger.TagRepo())
}

// CategoryUseCase implements UseCaseManager
func (u *useCaseManager) CategoryUseCase() usecase.CategoryUseCase {
	return usecase.NewCategoryUseCase(u.repoManger.CategoryRepo())
}

func (u *useCaseManager) UsersUseCase() usecase.UsersUseCase {
	return usecase.NewUsersUseCase(u.repoManger.UsersRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManger: repoManager}
}
