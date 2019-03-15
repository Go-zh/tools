package lsp

import (
	"context"
	"fmt"

	"github.com/Go-zh/tools/internal/lsp/protocol"
	"github.com/Go-zh/tools/internal/lsp/source"
	"github.com/Go-zh/tools/internal/span"
)

// formatRange formats a document with a given range.
func formatRange(ctx context.Context, v source.View, s span.Span) ([]protocol.TextEdit, error) {
	f, m, err := newColumnMap(ctx, v, s.URI)
	if err != nil {
		return nil, err
	}
	rng := s.Range(m.Converter)
	if rng.Start == rng.End {
		// If we have a single point, assume we want the whole file.
		tok := f.GetToken(ctx)
		if tok == nil {
			return nil, fmt.Errorf("no file information for %s", f.URI())
		}
		rng.End = tok.Pos(tok.Size())
	}
	edits, err := source.Format(ctx, f, rng)
	if err != nil {
		return nil, err
	}
	return toProtocolEdits(m, edits), nil
}

func toProtocolEdits(m *protocol.ColumnMapper, edits []source.TextEdit) []protocol.TextEdit {
	if edits == nil {
		return nil
	}
	result := make([]protocol.TextEdit, len(edits))
	for i, edit := range edits {
		result[i] = protocol.TextEdit{
			Range:   m.Range(edit.Span),
			NewText: edit.NewText,
		}
	}
	return result
}

func newColumnMap(ctx context.Context, v source.View, uri span.URI) (source.File, *protocol.ColumnMapper, error) {
	f, err := v.GetFile(ctx, uri)
	if err != nil {
		return nil, nil, err
	}
	tok := f.GetToken(ctx)
	if tok == nil {
		return nil, nil, fmt.Errorf("no file information for %v", f.URI())
	}
	m := protocol.NewColumnMapper(f.URI(), f.GetFileSet(ctx), tok, f.GetContent(ctx))
	return f, m, nil
}
