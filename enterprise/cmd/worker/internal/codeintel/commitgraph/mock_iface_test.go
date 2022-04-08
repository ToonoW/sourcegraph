// Code generated by go-mockgen 1.1.5; DO NOT EDIT.

package commitgraph

import (
	"context"
	"sync"
	"time"

	locker "github.com/sourcegraph/sourcegraph/internal/database/locker"
	gitserver "github.com/sourcegraph/sourcegraph/internal/gitserver"
	gitdomain "github.com/sourcegraph/sourcegraph/internal/gitserver/gitdomain"
)

// MockDBStore is a mock implementation of the DBStore interface (from the
// package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/worker/internal/codeintel/commitgraph)
// used for unit testing.
type MockDBStore struct {
	// CalculateVisibleUploadsFunc is an instance of a mock function object
	// controlling the behavior of the method CalculateVisibleUploads.
	CalculateVisibleUploadsFunc *DBStoreCalculateVisibleUploadsFunc
	// DirtyRepositoriesFunc is an instance of a mock function object
	// controlling the behavior of the method DirtyRepositories.
	DirtyRepositoriesFunc *DBStoreDirtyRepositoriesFunc
	// GetOldestCommitDateFunc is an instance of a mock function object
	// controlling the behavior of the method GetOldestCommitDate.
	GetOldestCommitDateFunc *DBStoreGetOldestCommitDateFunc
}

// NewMockDBStore creates a new mock of the DBStore interface. All methods
// return zero values for all results, unless overwritten.
func NewMockDBStore() *MockDBStore {
	return &MockDBStore{
		CalculateVisibleUploadsFunc: &DBStoreCalculateVisibleUploadsFunc{
			defaultHook: func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error {
				return nil
			},
		},
		DirtyRepositoriesFunc: &DBStoreDirtyRepositoriesFunc{
			defaultHook: func(context.Context) (map[int]int, error) {
				return nil, nil
			},
		},
		GetOldestCommitDateFunc: &DBStoreGetOldestCommitDateFunc{
			defaultHook: func(context.Context, int) (time.Time, bool, error) {
				return time.Time{}, false, nil
			},
		},
	}
}

// NewStrictMockDBStore creates a new mock of the DBStore interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockDBStore() *MockDBStore {
	return &MockDBStore{
		CalculateVisibleUploadsFunc: &DBStoreCalculateVisibleUploadsFunc{
			defaultHook: func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error {
				panic("unexpected invocation of MockDBStore.CalculateVisibleUploads")
			},
		},
		DirtyRepositoriesFunc: &DBStoreDirtyRepositoriesFunc{
			defaultHook: func(context.Context) (map[int]int, error) {
				panic("unexpected invocation of MockDBStore.DirtyRepositories")
			},
		},
		GetOldestCommitDateFunc: &DBStoreGetOldestCommitDateFunc{
			defaultHook: func(context.Context, int) (time.Time, bool, error) {
				panic("unexpected invocation of MockDBStore.GetOldestCommitDate")
			},
		},
	}
}

// NewMockDBStoreFrom creates a new mock of the MockDBStore interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockDBStoreFrom(i DBStore) *MockDBStore {
	return &MockDBStore{
		CalculateVisibleUploadsFunc: &DBStoreCalculateVisibleUploadsFunc{
			defaultHook: i.CalculateVisibleUploads,
		},
		DirtyRepositoriesFunc: &DBStoreDirtyRepositoriesFunc{
			defaultHook: i.DirtyRepositories,
		},
		GetOldestCommitDateFunc: &DBStoreGetOldestCommitDateFunc{
			defaultHook: i.GetOldestCommitDate,
		},
	}
}

// DBStoreCalculateVisibleUploadsFunc describes the behavior when the
// CalculateVisibleUploads method of the parent MockDBStore instance is
// invoked.
type DBStoreCalculateVisibleUploadsFunc struct {
	defaultHook func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error
	hooks       []func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error
	history     []DBStoreCalculateVisibleUploadsFuncCall
	mutex       sync.Mutex
}

// CalculateVisibleUploads delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockDBStore) CalculateVisibleUploads(v0 context.Context, v1 int, v2 *gitdomain.CommitGraph, v3 map[string][]gitdomain.RefDescription, v4 time.Duration, v5 time.Duration, v6 int) error {
	r0 := m.CalculateVisibleUploadsFunc.nextHook()(v0, v1, v2, v3, v4, v5, v6)
	m.CalculateVisibleUploadsFunc.appendCall(DBStoreCalculateVisibleUploadsFuncCall{v0, v1, v2, v3, v4, v5, v6, r0})
	return r0
}

// SetDefaultHook sets function that is called when the
// CalculateVisibleUploads method of the parent MockDBStore instance is
// invoked and the hook queue is empty.
func (f *DBStoreCalculateVisibleUploadsFunc) SetDefaultHook(hook func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// CalculateVisibleUploads method of the parent MockDBStore instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *DBStoreCalculateVisibleUploadsFunc) PushHook(hook func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *DBStoreCalculateVisibleUploadsFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *DBStoreCalculateVisibleUploadsFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error {
		return r0
	})
}

func (f *DBStoreCalculateVisibleUploadsFunc) nextHook() func(context.Context, int, *gitdomain.CommitGraph, map[string][]gitdomain.RefDescription, time.Duration, time.Duration, int) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *DBStoreCalculateVisibleUploadsFunc) appendCall(r0 DBStoreCalculateVisibleUploadsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of DBStoreCalculateVisibleUploadsFuncCall
// objects describing the invocations of this function.
func (f *DBStoreCalculateVisibleUploadsFunc) History() []DBStoreCalculateVisibleUploadsFuncCall {
	f.mutex.Lock()
	history := make([]DBStoreCalculateVisibleUploadsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// DBStoreCalculateVisibleUploadsFuncCall is an object that describes an
// invocation of method CalculateVisibleUploads on an instance of
// MockDBStore.
type DBStoreCalculateVisibleUploadsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 *gitdomain.CommitGraph
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 map[string][]gitdomain.RefDescription
	// Arg4 is the value of the 5th argument passed to this method
	// invocation.
	Arg4 time.Duration
	// Arg5 is the value of the 6th argument passed to this method
	// invocation.
	Arg5 time.Duration
	// Arg6 is the value of the 7th argument passed to this method
	// invocation.
	Arg6 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c DBStoreCalculateVisibleUploadsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3, c.Arg4, c.Arg5, c.Arg6}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c DBStoreCalculateVisibleUploadsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// DBStoreDirtyRepositoriesFunc describes the behavior when the
// DirtyRepositories method of the parent MockDBStore instance is invoked.
type DBStoreDirtyRepositoriesFunc struct {
	defaultHook func(context.Context) (map[int]int, error)
	hooks       []func(context.Context) (map[int]int, error)
	history     []DBStoreDirtyRepositoriesFuncCall
	mutex       sync.Mutex
}

// DirtyRepositories delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockDBStore) DirtyRepositories(v0 context.Context) (map[int]int, error) {
	r0, r1 := m.DirtyRepositoriesFunc.nextHook()(v0)
	m.DirtyRepositoriesFunc.appendCall(DBStoreDirtyRepositoriesFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the DirtyRepositories
// method of the parent MockDBStore instance is invoked and the hook queue
// is empty.
func (f *DBStoreDirtyRepositoriesFunc) SetDefaultHook(hook func(context.Context) (map[int]int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DirtyRepositories method of the parent MockDBStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *DBStoreDirtyRepositoriesFunc) PushHook(hook func(context.Context) (map[int]int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *DBStoreDirtyRepositoriesFunc) SetDefaultReturn(r0 map[int]int, r1 error) {
	f.SetDefaultHook(func(context.Context) (map[int]int, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *DBStoreDirtyRepositoriesFunc) PushReturn(r0 map[int]int, r1 error) {
	f.PushHook(func(context.Context) (map[int]int, error) {
		return r0, r1
	})
}

func (f *DBStoreDirtyRepositoriesFunc) nextHook() func(context.Context) (map[int]int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *DBStoreDirtyRepositoriesFunc) appendCall(r0 DBStoreDirtyRepositoriesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of DBStoreDirtyRepositoriesFuncCall objects
// describing the invocations of this function.
func (f *DBStoreDirtyRepositoriesFunc) History() []DBStoreDirtyRepositoriesFuncCall {
	f.mutex.Lock()
	history := make([]DBStoreDirtyRepositoriesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// DBStoreDirtyRepositoriesFuncCall is an object that describes an
// invocation of method DirtyRepositories on an instance of MockDBStore.
type DBStoreDirtyRepositoriesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 map[int]int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c DBStoreDirtyRepositoriesFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c DBStoreDirtyRepositoriesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// DBStoreGetOldestCommitDateFunc describes the behavior when the
// GetOldestCommitDate method of the parent MockDBStore instance is invoked.
type DBStoreGetOldestCommitDateFunc struct {
	defaultHook func(context.Context, int) (time.Time, bool, error)
	hooks       []func(context.Context, int) (time.Time, bool, error)
	history     []DBStoreGetOldestCommitDateFuncCall
	mutex       sync.Mutex
}

// GetOldestCommitDate delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockDBStore) GetOldestCommitDate(v0 context.Context, v1 int) (time.Time, bool, error) {
	r0, r1, r2 := m.GetOldestCommitDateFunc.nextHook()(v0, v1)
	m.GetOldestCommitDateFunc.appendCall(DBStoreGetOldestCommitDateFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the GetOldestCommitDate
// method of the parent MockDBStore instance is invoked and the hook queue
// is empty.
func (f *DBStoreGetOldestCommitDateFunc) SetDefaultHook(hook func(context.Context, int) (time.Time, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetOldestCommitDate method of the parent MockDBStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *DBStoreGetOldestCommitDateFunc) PushHook(hook func(context.Context, int) (time.Time, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *DBStoreGetOldestCommitDateFunc) SetDefaultReturn(r0 time.Time, r1 bool, r2 error) {
	f.SetDefaultHook(func(context.Context, int) (time.Time, bool, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *DBStoreGetOldestCommitDateFunc) PushReturn(r0 time.Time, r1 bool, r2 error) {
	f.PushHook(func(context.Context, int) (time.Time, bool, error) {
		return r0, r1, r2
	})
}

func (f *DBStoreGetOldestCommitDateFunc) nextHook() func(context.Context, int) (time.Time, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *DBStoreGetOldestCommitDateFunc) appendCall(r0 DBStoreGetOldestCommitDateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of DBStoreGetOldestCommitDateFuncCall objects
// describing the invocations of this function.
func (f *DBStoreGetOldestCommitDateFunc) History() []DBStoreGetOldestCommitDateFuncCall {
	f.mutex.Lock()
	history := make([]DBStoreGetOldestCommitDateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// DBStoreGetOldestCommitDateFuncCall is an object that describes an
// invocation of method GetOldestCommitDate on an instance of MockDBStore.
type DBStoreGetOldestCommitDateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 time.Time
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c DBStoreGetOldestCommitDateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c DBStoreGetOldestCommitDateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// MockGitserverClient is a mock implementation of the GitserverClient
// interface (from the package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/worker/internal/codeintel/commitgraph)
// used for unit testing.
type MockGitserverClient struct {
	// CommitGraphFunc is an instance of a mock function object controlling
	// the behavior of the method CommitGraph.
	CommitGraphFunc *GitserverClientCommitGraphFunc
	// RefDescriptionsFunc is an instance of a mock function object
	// controlling the behavior of the method RefDescriptions.
	RefDescriptionsFunc *GitserverClientRefDescriptionsFunc
}

// NewMockGitserverClient creates a new mock of the GitserverClient
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockGitserverClient() *MockGitserverClient {
	return &MockGitserverClient{
		CommitGraphFunc: &GitserverClientCommitGraphFunc{
			defaultHook: func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error) {
				return nil, nil
			},
		},
		RefDescriptionsFunc: &GitserverClientRefDescriptionsFunc{
			defaultHook: func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error) {
				return nil, nil
			},
		},
	}
}

// NewStrictMockGitserverClient creates a new mock of the GitserverClient
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockGitserverClient() *MockGitserverClient {
	return &MockGitserverClient{
		CommitGraphFunc: &GitserverClientCommitGraphFunc{
			defaultHook: func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error) {
				panic("unexpected invocation of MockGitserverClient.CommitGraph")
			},
		},
		RefDescriptionsFunc: &GitserverClientRefDescriptionsFunc{
			defaultHook: func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error) {
				panic("unexpected invocation of MockGitserverClient.RefDescriptions")
			},
		},
	}
}

// NewMockGitserverClientFrom creates a new mock of the MockGitserverClient
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockGitserverClientFrom(i GitserverClient) *MockGitserverClient {
	return &MockGitserverClient{
		CommitGraphFunc: &GitserverClientCommitGraphFunc{
			defaultHook: i.CommitGraph,
		},
		RefDescriptionsFunc: &GitserverClientRefDescriptionsFunc{
			defaultHook: i.RefDescriptions,
		},
	}
}

// GitserverClientCommitGraphFunc describes the behavior when the
// CommitGraph method of the parent MockGitserverClient instance is invoked.
type GitserverClientCommitGraphFunc struct {
	defaultHook func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error)
	hooks       []func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error)
	history     []GitserverClientCommitGraphFuncCall
	mutex       sync.Mutex
}

// CommitGraph delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockGitserverClient) CommitGraph(v0 context.Context, v1 int, v2 gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error) {
	r0, r1 := m.CommitGraphFunc.nextHook()(v0, v1, v2)
	m.CommitGraphFunc.appendCall(GitserverClientCommitGraphFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the CommitGraph method
// of the parent MockGitserverClient instance is invoked and the hook queue
// is empty.
func (f *GitserverClientCommitGraphFunc) SetDefaultHook(hook func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// CommitGraph method of the parent MockGitserverClient instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *GitserverClientCommitGraphFunc) PushHook(hook func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitserverClientCommitGraphFunc) SetDefaultReturn(r0 *gitdomain.CommitGraph, r1 error) {
	f.SetDefaultHook(func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitserverClientCommitGraphFunc) PushReturn(r0 *gitdomain.CommitGraph, r1 error) {
	f.PushHook(func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error) {
		return r0, r1
	})
}

func (f *GitserverClientCommitGraphFunc) nextHook() func(context.Context, int, gitserver.CommitGraphOptions) (*gitdomain.CommitGraph, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitserverClientCommitGraphFunc) appendCall(r0 GitserverClientCommitGraphFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitserverClientCommitGraphFuncCall objects
// describing the invocations of this function.
func (f *GitserverClientCommitGraphFunc) History() []GitserverClientCommitGraphFuncCall {
	f.mutex.Lock()
	history := make([]GitserverClientCommitGraphFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitserverClientCommitGraphFuncCall is an object that describes an
// invocation of method CommitGraph on an instance of MockGitserverClient.
type GitserverClientCommitGraphFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 gitserver.CommitGraphOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *gitdomain.CommitGraph
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c GitserverClientCommitGraphFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitserverClientCommitGraphFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// GitserverClientRefDescriptionsFunc describes the behavior when the
// RefDescriptions method of the parent MockGitserverClient instance is
// invoked.
type GitserverClientRefDescriptionsFunc struct {
	defaultHook func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error)
	hooks       []func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error)
	history     []GitserverClientRefDescriptionsFuncCall
	mutex       sync.Mutex
}

// RefDescriptions delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockGitserverClient) RefDescriptions(v0 context.Context, v1 int, v2 ...string) (map[string][]gitdomain.RefDescription, error) {
	r0, r1 := m.RefDescriptionsFunc.nextHook()(v0, v1, v2...)
	m.RefDescriptionsFunc.appendCall(GitserverClientRefDescriptionsFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the RefDescriptions
// method of the parent MockGitserverClient instance is invoked and the hook
// queue is empty.
func (f *GitserverClientRefDescriptionsFunc) SetDefaultHook(hook func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// RefDescriptions method of the parent MockGitserverClient instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *GitserverClientRefDescriptionsFunc) PushHook(hook func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *GitserverClientRefDescriptionsFunc) SetDefaultReturn(r0 map[string][]gitdomain.RefDescription, r1 error) {
	f.SetDefaultHook(func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *GitserverClientRefDescriptionsFunc) PushReturn(r0 map[string][]gitdomain.RefDescription, r1 error) {
	f.PushHook(func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error) {
		return r0, r1
	})
}

func (f *GitserverClientRefDescriptionsFunc) nextHook() func(context.Context, int, ...string) (map[string][]gitdomain.RefDescription, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *GitserverClientRefDescriptionsFunc) appendCall(r0 GitserverClientRefDescriptionsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of GitserverClientRefDescriptionsFuncCall
// objects describing the invocations of this function.
func (f *GitserverClientRefDescriptionsFunc) History() []GitserverClientRefDescriptionsFuncCall {
	f.mutex.Lock()
	history := make([]GitserverClientRefDescriptionsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// GitserverClientRefDescriptionsFuncCall is an object that describes an
// invocation of method RefDescriptions on an instance of
// MockGitserverClient.
type GitserverClientRefDescriptionsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg2 []string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 map[string][]gitdomain.RefDescription
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c GitserverClientRefDescriptionsFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg2 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0, c.Arg1}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c GitserverClientRefDescriptionsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// MockLocker is a mock implementation of the Locker interface (from the
// package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/worker/internal/codeintel/commitgraph)
// used for unit testing.
type MockLocker struct {
	// LockFunc is an instance of a mock function object controlling the
	// behavior of the method Lock.
	LockFunc *LockerLockFunc
}

// NewMockLocker creates a new mock of the Locker interface. All methods
// return zero values for all results, unless overwritten.
func NewMockLocker() *MockLocker {
	return &MockLocker{
		LockFunc: &LockerLockFunc{
			defaultHook: func(context.Context, int32, bool) (bool, locker.UnlockFunc, error) {
				return false, nil, nil
			},
		},
	}
}

// NewStrictMockLocker creates a new mock of the Locker interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockLocker() *MockLocker {
	return &MockLocker{
		LockFunc: &LockerLockFunc{
			defaultHook: func(context.Context, int32, bool) (bool, locker.UnlockFunc, error) {
				panic("unexpected invocation of MockLocker.Lock")
			},
		},
	}
}

// NewMockLockerFrom creates a new mock of the MockLocker interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockLockerFrom(i Locker) *MockLocker {
	return &MockLocker{
		LockFunc: &LockerLockFunc{
			defaultHook: i.Lock,
		},
	}
}

// LockerLockFunc describes the behavior when the Lock method of the parent
// MockLocker instance is invoked.
type LockerLockFunc struct {
	defaultHook func(context.Context, int32, bool) (bool, locker.UnlockFunc, error)
	hooks       []func(context.Context, int32, bool) (bool, locker.UnlockFunc, error)
	history     []LockerLockFuncCall
	mutex       sync.Mutex
}

// Lock delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockLocker) Lock(v0 context.Context, v1 int32, v2 bool) (bool, locker.UnlockFunc, error) {
	r0, r1, r2 := m.LockFunc.nextHook()(v0, v1, v2)
	m.LockFunc.appendCall(LockerLockFuncCall{v0, v1, v2, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the Lock method of the
// parent MockLocker instance is invoked and the hook queue is empty.
func (f *LockerLockFunc) SetDefaultHook(hook func(context.Context, int32, bool) (bool, locker.UnlockFunc, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Lock method of the parent MockLocker instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *LockerLockFunc) PushHook(hook func(context.Context, int32, bool) (bool, locker.UnlockFunc, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LockerLockFunc) SetDefaultReturn(r0 bool, r1 locker.UnlockFunc, r2 error) {
	f.SetDefaultHook(func(context.Context, int32, bool) (bool, locker.UnlockFunc, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LockerLockFunc) PushReturn(r0 bool, r1 locker.UnlockFunc, r2 error) {
	f.PushHook(func(context.Context, int32, bool) (bool, locker.UnlockFunc, error) {
		return r0, r1, r2
	})
}

func (f *LockerLockFunc) nextHook() func(context.Context, int32, bool) (bool, locker.UnlockFunc, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LockerLockFunc) appendCall(r0 LockerLockFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LockerLockFuncCall objects describing the
// invocations of this function.
func (f *LockerLockFunc) History() []LockerLockFuncCall {
	f.mutex.Lock()
	history := make([]LockerLockFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LockerLockFuncCall is an object that describes an invocation of method
// Lock on an instance of MockLocker.
type LockerLockFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int32
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 bool
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 bool
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 locker.UnlockFunc
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LockerLockFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LockerLockFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}
