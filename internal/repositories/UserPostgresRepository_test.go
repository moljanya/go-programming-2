package repositories

/*import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shifteducation/user-service/internal/entities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
func TestSaveUserOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()


	userService := mocks.NewMockUserService(mockCtrl)
	userService.EXPECT().
		GetById(gomock.Any(), id).
		Return(user, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/users/%s", id.String()), nil)

	router := NewRouter(userService)
	router.engine.ServeHTTP(w, req)

	}




/*
type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserPostgresRepository {
	return UserPostgresRepository{
		db: db,
	}
}

func (r UserPostgresRepository) Save(ctx context.Context, user entities.User) (*entities.User, error) {
	err := gorm.G[entities.User](r.db).Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserPostgresRepository) GetById(ctx context.Context, userId uuid.UUID) (*entities.User, error) {
	user, err := gorm.G[entities.User](r.db).
		Where("users.id = ?", userId).
		Joins(clause.LeftJoin.Association("Address"), nil).
		First(ctx)
	return &user, err
}

func (r UserPostgresRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	users, err := gorm.G[entities.User](r.db).
		Joins(clause.LeftJoin.Association("Address"), nil).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserPostgresRepository) Update(ctx context.Context, user entities.User) error {
	//log.Print("Not implemented")
	result, err := gorm.G[entities.User](r.db).Where("id = ?", user.Id).Updates(ctx, user)
	if err != nil {
		return err
	}

	if result == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r UserPostgresRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	//log.Print("Not implemented")
	result, err := gorm.G[entities.User](r.db).Where("id = ?", userId).Delete(ctx)
	if err != nil {
		return err
	}

	if result == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
*/
