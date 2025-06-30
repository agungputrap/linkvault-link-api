package tag

type Repository interface {
	ListTagsByUser(userID uint) ([]string, error)
	FindLinksByTag(userID uint, tag string) ([]uint, error)
}
