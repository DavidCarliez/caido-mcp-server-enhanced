package tools_test

// Tool list budgets in bytes, see TestToolListBudget. Measured when the
// budgets were set: default 91,978, compact 62,029, core 39,393, core with
// compact 23,315. The WebSocket draft tools add 3,691 bytes to default and
// 3,010 to core (31 tools).
const (
	budgetDefault     = 98000
	budgetCompact     = 66000
	budgetCore        = 43000
	budgetCoreCompact = 25000
)
