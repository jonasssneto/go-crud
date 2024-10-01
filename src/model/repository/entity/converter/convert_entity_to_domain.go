package converter

import (
	"github.com/jonasssneto/go-crud/src/model"
	"github.com/jonasssneto/go-crud/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.String())
	return domain
}
