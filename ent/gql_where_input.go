// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"time"

	"github.com/nibbleshift/drift/ent/predicate"
	"github.com/nibbleshift/drift/ent/user"
	"github.com/nibbleshift/drift/ent/utter"
)

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "username" field predicates.
	Username             *string  `json:"username,omitempty"`
	UsernameNEQ          *string  `json:"usernameNEQ,omitempty"`
	UsernameIn           []string `json:"usernameIn,omitempty"`
	UsernameNotIn        []string `json:"usernameNotIn,omitempty"`
	UsernameGT           *string  `json:"usernameGT,omitempty"`
	UsernameGTE          *string  `json:"usernameGTE,omitempty"`
	UsernameLT           *string  `json:"usernameLT,omitempty"`
	UsernameLTE          *string  `json:"usernameLTE,omitempty"`
	UsernameContains     *string  `json:"usernameContains,omitempty"`
	UsernameHasPrefix    *string  `json:"usernameHasPrefix,omitempty"`
	UsernameHasSuffix    *string  `json:"usernameHasSuffix,omitempty"`
	UsernameEqualFold    *string  `json:"usernameEqualFold,omitempty"`
	UsernameContainsFold *string  `json:"usernameContainsFold,omitempty"`

	// "first_name" field predicates.
	FirstName             *string  `json:"firstName,omitempty"`
	FirstNameNEQ          *string  `json:"firstNameNEQ,omitempty"`
	FirstNameIn           []string `json:"firstNameIn,omitempty"`
	FirstNameNotIn        []string `json:"firstNameNotIn,omitempty"`
	FirstNameGT           *string  `json:"firstNameGT,omitempty"`
	FirstNameGTE          *string  `json:"firstNameGTE,omitempty"`
	FirstNameLT           *string  `json:"firstNameLT,omitempty"`
	FirstNameLTE          *string  `json:"firstNameLTE,omitempty"`
	FirstNameContains     *string  `json:"firstNameContains,omitempty"`
	FirstNameHasPrefix    *string  `json:"firstNameHasPrefix,omitempty"`
	FirstNameHasSuffix    *string  `json:"firstNameHasSuffix,omitempty"`
	FirstNameEqualFold    *string  `json:"firstNameEqualFold,omitempty"`
	FirstNameContainsFold *string  `json:"firstNameContainsFold,omitempty"`

	// "last_name" field predicates.
	LastName             *string  `json:"lastName,omitempty"`
	LastNameNEQ          *string  `json:"lastNameNEQ,omitempty"`
	LastNameIn           []string `json:"lastNameIn,omitempty"`
	LastNameNotIn        []string `json:"lastNameNotIn,omitempty"`
	LastNameGT           *string  `json:"lastNameGT,omitempty"`
	LastNameGTE          *string  `json:"lastNameGTE,omitempty"`
	LastNameLT           *string  `json:"lastNameLT,omitempty"`
	LastNameLTE          *string  `json:"lastNameLTE,omitempty"`
	LastNameContains     *string  `json:"lastNameContains,omitempty"`
	LastNameHasPrefix    *string  `json:"lastNameHasPrefix,omitempty"`
	LastNameHasSuffix    *string  `json:"lastNameHasSuffix,omitempty"`
	LastNameEqualFold    *string  `json:"lastNameEqualFold,omitempty"`
	LastNameContainsFold *string  `json:"lastNameContainsFold,omitempty"`

	// "utters" edge predicates.
	HasUtters     *bool              `json:"hasUtters,omitempty"`
	HasUttersWith []*UtterWhereInput `json:"hasUttersWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Username != nil {
		predicates = append(predicates, user.UsernameEQ(*i.Username))
	}
	if i.UsernameNEQ != nil {
		predicates = append(predicates, user.UsernameNEQ(*i.UsernameNEQ))
	}
	if len(i.UsernameIn) > 0 {
		predicates = append(predicates, user.UsernameIn(i.UsernameIn...))
	}
	if len(i.UsernameNotIn) > 0 {
		predicates = append(predicates, user.UsernameNotIn(i.UsernameNotIn...))
	}
	if i.UsernameGT != nil {
		predicates = append(predicates, user.UsernameGT(*i.UsernameGT))
	}
	if i.UsernameGTE != nil {
		predicates = append(predicates, user.UsernameGTE(*i.UsernameGTE))
	}
	if i.UsernameLT != nil {
		predicates = append(predicates, user.UsernameLT(*i.UsernameLT))
	}
	if i.UsernameLTE != nil {
		predicates = append(predicates, user.UsernameLTE(*i.UsernameLTE))
	}
	if i.UsernameContains != nil {
		predicates = append(predicates, user.UsernameContains(*i.UsernameContains))
	}
	if i.UsernameHasPrefix != nil {
		predicates = append(predicates, user.UsernameHasPrefix(*i.UsernameHasPrefix))
	}
	if i.UsernameHasSuffix != nil {
		predicates = append(predicates, user.UsernameHasSuffix(*i.UsernameHasSuffix))
	}
	if i.UsernameEqualFold != nil {
		predicates = append(predicates, user.UsernameEqualFold(*i.UsernameEqualFold))
	}
	if i.UsernameContainsFold != nil {
		predicates = append(predicates, user.UsernameContainsFold(*i.UsernameContainsFold))
	}
	if i.FirstName != nil {
		predicates = append(predicates, user.FirstNameEQ(*i.FirstName))
	}
	if i.FirstNameNEQ != nil {
		predicates = append(predicates, user.FirstNameNEQ(*i.FirstNameNEQ))
	}
	if len(i.FirstNameIn) > 0 {
		predicates = append(predicates, user.FirstNameIn(i.FirstNameIn...))
	}
	if len(i.FirstNameNotIn) > 0 {
		predicates = append(predicates, user.FirstNameNotIn(i.FirstNameNotIn...))
	}
	if i.FirstNameGT != nil {
		predicates = append(predicates, user.FirstNameGT(*i.FirstNameGT))
	}
	if i.FirstNameGTE != nil {
		predicates = append(predicates, user.FirstNameGTE(*i.FirstNameGTE))
	}
	if i.FirstNameLT != nil {
		predicates = append(predicates, user.FirstNameLT(*i.FirstNameLT))
	}
	if i.FirstNameLTE != nil {
		predicates = append(predicates, user.FirstNameLTE(*i.FirstNameLTE))
	}
	if i.FirstNameContains != nil {
		predicates = append(predicates, user.FirstNameContains(*i.FirstNameContains))
	}
	if i.FirstNameHasPrefix != nil {
		predicates = append(predicates, user.FirstNameHasPrefix(*i.FirstNameHasPrefix))
	}
	if i.FirstNameHasSuffix != nil {
		predicates = append(predicates, user.FirstNameHasSuffix(*i.FirstNameHasSuffix))
	}
	if i.FirstNameEqualFold != nil {
		predicates = append(predicates, user.FirstNameEqualFold(*i.FirstNameEqualFold))
	}
	if i.FirstNameContainsFold != nil {
		predicates = append(predicates, user.FirstNameContainsFold(*i.FirstNameContainsFold))
	}
	if i.LastName != nil {
		predicates = append(predicates, user.LastNameEQ(*i.LastName))
	}
	if i.LastNameNEQ != nil {
		predicates = append(predicates, user.LastNameNEQ(*i.LastNameNEQ))
	}
	if len(i.LastNameIn) > 0 {
		predicates = append(predicates, user.LastNameIn(i.LastNameIn...))
	}
	if len(i.LastNameNotIn) > 0 {
		predicates = append(predicates, user.LastNameNotIn(i.LastNameNotIn...))
	}
	if i.LastNameGT != nil {
		predicates = append(predicates, user.LastNameGT(*i.LastNameGT))
	}
	if i.LastNameGTE != nil {
		predicates = append(predicates, user.LastNameGTE(*i.LastNameGTE))
	}
	if i.LastNameLT != nil {
		predicates = append(predicates, user.LastNameLT(*i.LastNameLT))
	}
	if i.LastNameLTE != nil {
		predicates = append(predicates, user.LastNameLTE(*i.LastNameLTE))
	}
	if i.LastNameContains != nil {
		predicates = append(predicates, user.LastNameContains(*i.LastNameContains))
	}
	if i.LastNameHasPrefix != nil {
		predicates = append(predicates, user.LastNameHasPrefix(*i.LastNameHasPrefix))
	}
	if i.LastNameHasSuffix != nil {
		predicates = append(predicates, user.LastNameHasSuffix(*i.LastNameHasSuffix))
	}
	if i.LastNameEqualFold != nil {
		predicates = append(predicates, user.LastNameEqualFold(*i.LastNameEqualFold))
	}
	if i.LastNameContainsFold != nil {
		predicates = append(predicates, user.LastNameContainsFold(*i.LastNameContainsFold))
	}

	if i.HasUtters != nil {
		p := user.HasUtters()
		if !*i.HasUtters {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUttersWith) > 0 {
		with := make([]predicate.Utter, 0, len(i.HasUttersWith))
		for _, w := range i.HasUttersWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUttersWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasUttersWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}

// UtterWhereInput represents a where input for filtering Utter queries.
type UtterWhereInput struct {
	Predicates []predicate.Utter  `json:"-"`
	Not        *UtterWhereInput   `json:"not,omitempty"`
	Or         []*UtterWhereInput `json:"or,omitempty"`
	And        []*UtterWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "data" field predicates.
	Data             *string  `json:"data,omitempty"`
	DataNEQ          *string  `json:"dataNEQ,omitempty"`
	DataIn           []string `json:"dataIn,omitempty"`
	DataNotIn        []string `json:"dataNotIn,omitempty"`
	DataGT           *string  `json:"dataGT,omitempty"`
	DataGTE          *string  `json:"dataGTE,omitempty"`
	DataLT           *string  `json:"dataLT,omitempty"`
	DataLTE          *string  `json:"dataLTE,omitempty"`
	DataContains     *string  `json:"dataContains,omitempty"`
	DataHasPrefix    *string  `json:"dataHasPrefix,omitempty"`
	DataHasSuffix    *string  `json:"dataHasSuffix,omitempty"`
	DataEqualFold    *string  `json:"dataEqualFold,omitempty"`
	DataContainsFold *string  `json:"dataContainsFold,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UtterWhereInput) AddPredicates(predicates ...predicate.Utter) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UtterWhereInput filter on the UtterQuery builder.
func (i *UtterWhereInput) Filter(q *UtterQuery) (*UtterQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUtterWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUtterWhereInput is returned in case the UtterWhereInput is empty.
var ErrEmptyUtterWhereInput = errors.New("ent: empty predicate UtterWhereInput")

// P returns a predicate for filtering utters.
// An error is returned if the input is empty or invalid.
func (i *UtterWhereInput) P() (predicate.Utter, error) {
	var predicates []predicate.Utter
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, utter.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Utter, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, utter.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Utter, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, utter.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, utter.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, utter.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, utter.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, utter.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, utter.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, utter.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, utter.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, utter.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, utter.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, utter.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, utter.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, utter.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, utter.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, utter.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, utter.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, utter.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.Data != nil {
		predicates = append(predicates, utter.DataEQ(*i.Data))
	}
	if i.DataNEQ != nil {
		predicates = append(predicates, utter.DataNEQ(*i.DataNEQ))
	}
	if len(i.DataIn) > 0 {
		predicates = append(predicates, utter.DataIn(i.DataIn...))
	}
	if len(i.DataNotIn) > 0 {
		predicates = append(predicates, utter.DataNotIn(i.DataNotIn...))
	}
	if i.DataGT != nil {
		predicates = append(predicates, utter.DataGT(*i.DataGT))
	}
	if i.DataGTE != nil {
		predicates = append(predicates, utter.DataGTE(*i.DataGTE))
	}
	if i.DataLT != nil {
		predicates = append(predicates, utter.DataLT(*i.DataLT))
	}
	if i.DataLTE != nil {
		predicates = append(predicates, utter.DataLTE(*i.DataLTE))
	}
	if i.DataContains != nil {
		predicates = append(predicates, utter.DataContains(*i.DataContains))
	}
	if i.DataHasPrefix != nil {
		predicates = append(predicates, utter.DataHasPrefix(*i.DataHasPrefix))
	}
	if i.DataHasSuffix != nil {
		predicates = append(predicates, utter.DataHasSuffix(*i.DataHasSuffix))
	}
	if i.DataEqualFold != nil {
		predicates = append(predicates, utter.DataEqualFold(*i.DataEqualFold))
	}
	if i.DataContainsFold != nil {
		predicates = append(predicates, utter.DataContainsFold(*i.DataContainsFold))
	}

	if i.HasOwner != nil {
		p := utter.HasOwner()
		if !*i.HasOwner {
			p = utter.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOwnerWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, utter.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUtterWhereInput
	case 1:
		return predicates[0], nil
	default:
		return utter.And(predicates...), nil
	}
}
