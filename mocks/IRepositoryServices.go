// Code generated by mockery v1.0.0
package mocks

import context "context"
import github "github.com/google/go-github/github"
import mock "github.com/stretchr/testify/mock"

// IRepositoryServices is an autogenerated mock type for the IRepositoryServices type
type IRepositoryServices struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, owner, repo
func (_m *IRepositoryServices) Get(ctx context.Context, owner string, repo string) (*github.Repository, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 *github.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *github.Repository); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Repository)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, owner, repo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetReadme provides a mock function with given fields: ctx, owner, repo, opt
func (_m *IRepositoryServices) GetReadme(ctx context.Context, owner string, repo string, opt *github.RepositoryContentGetOptions) (*github.RepositoryContent, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, opt)

	var r0 *github.RepositoryContent
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.RepositoryContentGetOptions) *github.RepositoryContent); ok {
		r0 = rf(ctx, owner, repo, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryContent)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.RepositoryContentGetOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.RepositoryContentGetOptions) error); ok {
		r2 = rf(ctx, owner, repo, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListCommits provides a mock function with given fields: ctx, owner, repo, opt
func (_m *IRepositoryServices) ListCommits(ctx context.Context, owner string, repo string, opt *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, opt)

	var r0 []*github.RepositoryCommit
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.CommitsListOptions) []*github.RepositoryCommit); ok {
		r0 = rf(ctx, owner, repo, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.RepositoryCommit)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.CommitsListOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.CommitsListOptions) error); ok {
		r2 = rf(ctx, owner, repo, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
