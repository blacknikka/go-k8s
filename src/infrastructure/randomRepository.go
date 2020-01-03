package infrastructure

import (
	"github.com/blacknikka/go-k8s/domain"
	"github.com/gomodule/redigo/redis"
)

type randomRepository struct {
	Conn redis.Conn
}

type RandomRepository interface {
	Save(random domain.RandomItem) error
	Update(random domain.RandomItem) error
	Delete(random domain.RandomItem) error
	Find(name string) string
}

func NewRandomRepository(conn redis.Conn) *randomRepository {
	return &randomRepository{Conn: conn}
}

func (r *randomRepository) Save(random domain.RandomItem) error {
	_, err := r.Conn.Do("SET", random.Name, random.Random)
	return err
}

func (r *randomRepository) Update(random domain.RandomItem) error {
	return nil
}

func (r *randomRepository) Delete(random domain.RandomItem) error {
	return nil
}

func (r *randomRepository) Find(random domain.RandomItem) string {
	return "aaa"
}
