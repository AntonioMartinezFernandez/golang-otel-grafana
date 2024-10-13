package rolldice

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	// "go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const name = "github.com/AntonioMartinezFernandez/golang-otel-grafana"

func NewRollDiceHandler() http.HandlerFunc {
	tracer := otel.Tracer(name)
	meter := otel.Meter(name)

	rollCnt, err := meter.Int64Counter("dice.rolls",
		metric.WithDescription("The number of rolls by roll value"),
		metric.WithUnit("{roll}"))
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "roll")
		defer span.End()

		roll := 1 + rand.Intn(6) //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand) is ignored as this is not security-sensitive.

		var msg string
		if player := r.PathValue("player"); player != "" {
			msg = fmt.Sprintf("%s is rolling the dice", player)
		} else {
			msg = "Anonymous player is rolling the dice"
		}
		// logger.InfoContext(ctx, msg, "result", roll)
		fmt.Println(msg)

		rollValueAttr := attribute.Int("roll.value", roll)
		span.SetAttributes(rollValueAttr)
		rollCnt.Add(ctx, 1, metric.WithAttributes(rollValueAttr))

		resp := strconv.Itoa(roll) + "\n"
		if _, err := io.WriteString(w, resp); err != nil {
			// logger.ErrorContext(ctx, "Write failed", "error", err)
			fmt.Println(err)
		}
	}
}
