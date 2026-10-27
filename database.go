package main

import (
	"context"
	"time"
)

type Storage interface {
	InsertRequestLog(context.Context, int, string, error, time.Time)
}

type DBService struct {
	next    Service
	storage Storage
}

func NewDBService(storage Storage, next Service) Service {
	return &DBService{
		storage: storage,
		next:    next,
	}
}

func (s *DBService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(begin time.Time) {
		s.storage.InsertRequestLog(ctx, ctx.Value("requestID").(int), fact.Fact, err, begin)
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
