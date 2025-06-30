package link

type Repository interface {
	Create(link *Link) error
	FindByUser(userID uint) ([]Link, error)
	Delete(id uint, userID uint) error
	Update(link *Link) error
}
