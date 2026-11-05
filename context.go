package logger

import "context"

// AddContextField returns a context with a key/value field added to it
func AddContextField(ctx context.Context, key, value string) context.Context {
	if existingMap, ok := ctx.Value(contextLoggingFieldMapKey).(map[string]string); ok {
		existingMap[key] = value
		return ctx
	}

	return context.WithValue(ctx, contextLoggingFieldMapKey, map[string]string{key: value})
}
