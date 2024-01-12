package dependency

import (
	"github.com/alexisleon/stori/internal/mailer"
	"github.com/alexisleon/stori/internal/storage"
)

type SettlementHandler struct {
	db     *storage.Conn
	mailer mailer.Mailer
}

func NewSettlementHandler(c *storage.Conn, m mailer.Mailer) *SettlementHandler {
	return &SettlementHandler{
		db:     c,
		mailer: m,
	}
}

func (h *SettlementHandler) ProcessSettlement() error {
	return nil
}
