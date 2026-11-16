package tools_test

// Tool list budgets in bytes, see TestToolListBudget. Measured when the
// budgets were set: default 91,978, compact 62,029, core 39,393, core with
// compact 23,315. This fork adds four tools (plugin call, WebSocket draft
// write/read and clear) on top of upstream: default 105,978, compact 71,155,
// core 45,393, core with compact 26,834.
const (
	budgetDefault     = 108000
	budgetCompact     = 73000
	budgetCore        = 47000
	budgetCoreCompact = 28000
)
