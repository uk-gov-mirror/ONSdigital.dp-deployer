package deployment

import "fmt"

// AllocationError is an error implemtation that includes pending, runnning and total
// allocation counts.
type AllocationError struct {
	pending int
	running int
	total   int
}

func (e *AllocationError) Error() string {
	return fmt.Sprintf("failed to start all allocations (pending: %d, running: %d, total: %d)", e.pending, e.running, e.total)
}

// ClientResponseError is an error implemtation that includes the body and status
// code of the response.
type ClientResponseError struct {
	body       string
	statusCode int
}

func (e *ClientResponseError) Error() string {
	return fmt.Sprintf("unexpected response from client (%d): %s", e.statusCode, e.body)
}

// EvaluationError is an error implemtation that includes the id of the evaluation.
type EvaluationError struct {
	id string
}

func (e *EvaluationError) Error() string {
	return fmt.Sprintf("error occurred for evaluation: %s", e.id)
}

// PlanError is an error implemtation that includes the errors or warnings
type PlanError struct {
	errors   string
	service  string
	warnings string
}

func (e *PlanError) Error() string {
	return fmt.Sprintf("plan for service %s generated errors or warnings (errors: %s, warnings: %s)", e.service, e.errors, e.warnings)
}

// TimeoutError is an error implemtation that includes the action that timed out.
type TimeoutError struct {
	action string
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("timed out waiting for action to complete: %s", e.action)
}