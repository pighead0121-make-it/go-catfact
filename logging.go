package main

import (
	"context"
	"log/slog"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(begin time.Time) {
		slog.InfoContext(ctx, "getfact",
			"requestID", ctx.Value("requestID"),
			"fact", fact.Fact,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
