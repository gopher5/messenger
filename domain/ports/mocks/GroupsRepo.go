// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"
	errors "messanger/pkg/errors"

	mock "github.com/stretchr/testify/mock"

	models "messanger/domain/models"
)

// GroupsRepo is an autogenerated mock type for the GroupsRepo type
type GroupsRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *GroupsRepo) Delete(ctx context.Context, id int) *errors.Error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, int) *errors.Error); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.Error)
		}
	}

	return r0
}

// GetById provides a mock function with given fields: ctx, id
func (_m *GroupsRepo) GetById(ctx context.Context, id int) (*models.Group, *errors.Error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *models.Group
	var r1 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.Group, *errors.Error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Group); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) *errors.Error); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.Error)
		}
	}

	return r0, r1
}

// GetGroupByChatId provides a mock function with given fields: ctx, chatId
func (_m *GroupsRepo) GetGroupByChatId(ctx context.Context, chatId int) (*models.Group, *errors.Error) {
	ret := _m.Called(ctx, chatId)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupByChatId")
	}

	var r0 *models.Group
	var r1 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.Group, *errors.Error)); ok {
		return rf(ctx, chatId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Group); ok {
		r0 = rf(ctx, chatId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) *errors.Error); ok {
		r1 = rf(ctx, chatId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.Error)
		}
	}

	return r0, r1
}

// GetGroupsByUser provides a mock function with given fields: ctx, userId
func (_m *GroupsRepo) GetGroupsByUser(ctx context.Context, userId int) ([]models.Group, *errors.Error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupsByUser")
	}

	var r0 []models.Group
	var r1 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]models.Group, *errors.Error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []models.Group); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) *errors.Error); ok {
		r1 = rf(ctx, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.Error)
		}
	}

	return r0, r1
}

// GetRole provides a mock function with given fields: ctx, userId, groupId
func (_m *GroupsRepo) GetRole(ctx context.Context, userId int, groupId int) (string, *errors.Error) {
	ret := _m.Called(ctx, userId, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetRole")
	}

	var r0 string
	var r1 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (string, *errors.Error)); ok {
		return rf(ctx, userId, groupId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) string); ok {
		r0 = rf(ctx, userId, groupId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) *errors.Error); ok {
		r1 = rf(ctx, userId, groupId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.Error)
		}
	}

	return r0, r1
}

// GetUsersByGroup provides a mock function with given fields: ctx, id
func (_m *GroupsRepo) GetUsersByGroup(ctx context.Context, id int) ([]int, *errors.Error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUsersByGroup")
	}

	var r0 []int
	var r1 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]int, *errors.Error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []int); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) *errors.Error); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.Error)
		}
	}

	return r0, r1
}

// New provides a mock function with given fields: ctx, group
func (_m *GroupsRepo) New(ctx context.Context, group *models.Group) *errors.Error {
	ret := _m.Called(ctx, group)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Group) *errors.Error); ok {
		r0 = rf(ctx, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.Error)
		}
	}

	return r0
}

// SetRole provides a mock function with given fields: ctx, userId, groupId, role
func (_m *GroupsRepo) SetRole(ctx context.Context, userId int, groupId int, role string) *errors.Error {
	ret := _m.Called(ctx, userId, groupId, role)

	if len(ret) == 0 {
		panic("no return value specified for SetRole")
	}

	var r0 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) *errors.Error); ok {
		r0 = rf(ctx, userId, groupId, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.Error)
		}
	}

	return r0
}

// Update provides a mock function with given fields: ctx, group
func (_m *GroupsRepo) Update(ctx context.Context, group *models.Group) *errors.Error {
	ret := _m.Called(ctx, group)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Group) *errors.Error); ok {
		r0 = rf(ctx, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.Error)
		}
	}

	return r0
}

// NewGroupsRepo creates a new instance of GroupsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupsRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupsRepo {
	mock := &GroupsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
