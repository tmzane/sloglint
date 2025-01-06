package mixed_args

import (
	"context"
	"log/slog"
)

func tests() {
	logger := slog.New(nil)
	ctx := context.Background()

	slog.Info("msg")
	slog.Info("msg", "foo", 1)
	slog.Info("msg", "foo", 1, "bar", 2)
	slog.Info("msg", slog.Int("foo", 1))
	slog.Info("msg", slog.Int("foo", 1), slog.Int("bar", 2))
	slog.With("foo", 1, "bar", 2).Info("msg")
	slog.With(slog.Int("foo", 1)).Info("msg")
	slog.With(slog.Int("foo", 1), slog.Int("bar", 2)).Info("msg")
	slog.Group("group", "foo", 1, "baz", 2)
	slog.Group("group", slog.String("foo", "1"))

	slog.Log(ctx, slog.LevelInfo, "msg", "foo", 1, slog.Int("bar", 2)) // want `key-value pairs and attributes should not be mixed`
	slog.Debug("msg", "foo", 1, slog.Int("bar", 2))                    // want `key-value pairs and attributes should not be mixed`
	slog.Info("msg", "foo", 1, slog.Int("bar", 2))                     // want `key-value pairs and attributes should not be mixed`
	slog.Warn("msg", "foo", 1, slog.Int("bar", 2))                     // want `key-value pairs and attributes should not be mixed`
	slog.Error("msg", "foo", 1, slog.Int("bar", 2))                    // want `key-value pairs and attributes should not be mixed`
	slog.DebugContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))        // want `key-value pairs and attributes should not be mixed`
	slog.InfoContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))         // want `key-value pairs and attributes should not be mixed`
	slog.WarnContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))         // want `key-value pairs and attributes should not be mixed`
	slog.ErrorContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))        // want `key-value pairs and attributes should not be mixed`
	slog.With("foo", 1, slog.Int("bar", 2)).ErrorContext(ctx, "msg")   // want `key-value pairs and attributes should not be mixed`
	slog.Group("group", slog.String("foo", "1"), "baz", 2)             // want `key-value pairs and attributes should not be mixed`

	logger.Log(ctx, slog.LevelInfo, "msg", "foo", 1, slog.Int("bar", 2))         // want `key-value pairs and attributes should not be mixed`
	logger.Debug("msg", "foo", 1, slog.Int("bar", 2))                            // want `key-value pairs and attributes should not be mixed`
	logger.Info("msg", "foo", 1, slog.Int("bar", 2))                             // want `key-value pairs and attributes should not be mixed`
	logger.Warn("msg", "foo", 1, slog.Int("bar", 2))                             // want `key-value pairs and attributes should not be mixed`
	logger.Error("msg", "foo", 1, slog.Int("bar", 2))                            // want `key-value pairs and attributes should not be mixed`
	logger.DebugContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))                // want `key-value pairs and attributes should not be mixed`
	logger.InfoContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))                 // want `key-value pairs and attributes should not be mixed`
	logger.WarnContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))                 // want `key-value pairs and attributes should not be mixed`
	logger.ErrorContext(ctx, "msg", "foo", 1, slog.Int("bar", 2))                // want `key-value pairs and attributes should not be mixed`
	logger.With("foo", 1, slog.Int("bar", 2)).ErrorContext(ctx, "msg")           // want `key-value pairs and attributes should not be mixed`
	logger.With("foo", 1).ErrorContext(ctx, "msg", "foo", 1, slog.Int("bar", 2)) // want `key-value pairs and attributes should not be mixed`
}
