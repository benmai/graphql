package graphql_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"sync"
	"testing"

	"github.com/matryer/is"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/location"
	"github.com/graphql-go/graphql/testutil"
)

func tinit(t *testing.T) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Type",
			Fields: graphql.Fields{
				"a": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "foo", nil
					},
				},
			},
		}),
	})
	if err != nil {
		t.Fatalf("Error in schema %v", err.Error())
	}
	return schema
}

func TestExtensionInitPanic(t *testing.T) {
	ext := newTestExt()
	ext.initFn = func(ctx context.Context, p *graphql.Params) context.Context {
		if true {
			panic(errors.New("test error"))
		}
		return ctx
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.Init: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionParseDidStartPanic(t *testing.T) {
	ext := newTestExt()
	ext.parseDidStartFn = func(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func(err error) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ParseDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionParseFinishFuncPanic(t *testing.T) {
	ext := newTestExt()
	ext.parseDidStartFn = func(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
		return ctx, func(err error) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ParseFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionValidationDidStartPanic(t *testing.T) {
	ext := newTestExt()
	ext.validationDidStartFn = func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func([]gqlerrors.FormattedError) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ValidationDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionValidationFinishFuncPanic(t *testing.T) {
	ext := newTestExt()
	ext.validationDidStartFn = func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
		return ctx, func([]gqlerrors.FormattedError) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ValidationFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionExecutionDidStartPanic(t *testing.T) {
	ext := newTestExt()
	ext.executionDidStartFn = func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func(r *graphql.Result) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: nil,
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ExecutionDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionExecutionFinishFuncPanic(t *testing.T) {
	ext := newTestExt()
	ext.executionDidStartFn = func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
		return ctx, func(r *graphql.Result) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ExecutionFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionResolveFieldDidStartPanic(t *testing.T) {
	ext := newTestExt()
	ext.resolveFieldDidStartFn = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
		if true {
			panic(errors.New("test error"))
		}
		return ctx, func(v interface{}, err error) {

		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ResolveFieldDidStart: %v", ext.Name(), errors.New("test error"))),
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionResolveFieldFinish(t *testing.T) {
	ext := newExtensionMock()

	type resolveFieldFinishCall struct {
		fieldName string
		value     interface{}
		err       error
	}
	var mockResolveFieldFinishFuncCalls []resolveFieldFinishCall
	lockMockResolveFieldFinishFuncCalls := &sync.Mutex{}

	ext.ResolveFieldDidStartFunc = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
		return ctx, func(value interface{}, err error) {
			lockMockResolveFieldFinishFuncCalls.Lock()
			mockResolveFieldFinishFuncCalls = append(
				mockResolveFieldFinishFuncCalls,
				resolveFieldFinishCall{i.FieldName, value, err},
			)
			lockMockResolveFieldFinishFuncCalls.Unlock()
		}

	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Type",
			Fields: graphql.Fields{
				"asyncValue": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return func() (interface{}, error) {
							return "asyncValue value", nil
						}, nil
					},
				},
				"syncValue": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "syncValue value", nil
					},
				},
				"asyncValueError": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return func() (interface{}, error) {
							return nil, errors.New("asyncValueError error")
						}, nil
					},
				},
				"syncValueError": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return nil, errors.New("syncValueError error")
					},
				},
			},
		}),
	})
	if err != nil {
		t.Fatalf("Error in schema %v", err.Error())
	}

	query := `query Example {
		syncValue
		asyncValue
		syncValueError
		asyncValueError
	}`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	syncValueError := gqlerrors.NewFormattedError("syncValueError error")
	syncValueError.Message = "syncValueError error"
	syncValueError.Locations = []location.SourceLocation{
		{
			Line:   4,
			Column: 3,
		},
	}
	syncValueError.Path = []interface{}{"syncValueError"}

	asyncValueError := gqlerrors.NewFormattedError("asyncValueError error")
	asyncValueError.Message = "asyncValueError error"
	asyncValueError.Locations = []location.SourceLocation{
		{
			Line:   5,
			Column: 3,
		},
	}
	asyncValueError.Path = []interface{}{"asyncValueError"}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"syncValue":       "syncValue value",
			"asyncValue":      "asyncValue value",
			"syncValueError":  nil,
			"asyncValueError": nil,
		},
		Errors: []gqlerrors.FormattedError{
			syncValueError,
			asyncValueError,
		},
	}

	// Check value return.
	is := is.NewRelaxed(t)
	is.Equal(expected.Data, result.Data)
	is.True(testutil.EqualFormattedErrors(expected.Errors, result.Errors))

	// Check that ResolveFieldDidStart was called for each field.
	calls := ext.ResolveFieldDidStartCalls()
	is.Equal(4, len(calls))
	for _, fieldName := range []string{
		"syncValueError",
		"asyncValueError",
		"syncValue",
		"asyncValue",
	} {
		var found bool
		for _, call := range calls {
			if call.In2.FieldName == fieldName {
				found = true
			}
		}
		is.True(found)
	}

	// Check that ResolveFieldFinishFunc was called for each field.
	sort.Slice(mockResolveFieldFinishFuncCalls, func(i, j int) bool {
		return mockResolveFieldFinishFuncCalls[i].fieldName < mockResolveFieldFinishFuncCalls[j].fieldName
	})

	expectedCalls := []resolveFieldFinishCall{
		{
			fieldName: "asyncValue",
			value:     "asyncValue value",
		},
		{
			fieldName: "asyncValueError",
			err:       errors.New("asyncValueError error"),
		},
		{
			fieldName: "syncValue",
			value:     "syncValue value",
		},
		{
			fieldName: "syncValueError",
			err:       errors.New("syncValueError error"),
		},
	}

	is.Equal(len(expectedCalls), len(mockResolveFieldFinishFuncCalls))
	for i, c := range mockResolveFieldFinishFuncCalls {
		expectedCall := expectedCalls[i]
		is.Equal(expectedCall, c)
	}
}

func TestExtensionResolveFieldFinishFuncPanic(t *testing.T) {
	ext := newTestExt()
	ext.resolveFieldDidStartFn = func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
		return ctx, func(v interface{}, err error) {
			panic(errors.New("test error"))
		}
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.ResolveFieldFinishFunc: %v", ext.Name(), errors.New("test error"))),
		},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func TestExtensionGetResultPanic(t *testing.T) {
	ext := newTestExt()
	ext.getResultFn = func(context.Context) interface{} {
		if true {
			panic(errors.New("test error"))
		}
		return nil
	}
	ext.hasResultFn = func() bool {
		return true
	}

	schema := tinit(t)
	query := `query Example { a }`
	schema.AddExtensions(ext)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"a": "foo",
		},
		Errors: []gqlerrors.FormattedError{
			gqlerrors.FormatError(fmt.Errorf("%s.GetResult: %v", ext.Name(), errors.New("test error"))),
		},
		Extensions: make(map[string]interface{}),
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Unexpected result, Diff: %v", testutil.Diff(expected, result))
	}
}

func newExtensionMock() *graphql.ExtensionMock {
	return &graphql.ExtensionMock{
		InitFunc: func(ctx context.Context, p *graphql.Params) context.Context {
			return ctx
		},
		NameFunc: func() string {
			return "mockExtension"
		},
		ParseDidStartFunc: func(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
			return ctx, func(error) {}
		},
		ValidationDidStartFunc: func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
			return ctx, func([]gqlerrors.FormattedError) {}
		},
		ExecutionDidStartFunc: func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
			return ctx, func(*graphql.Result) {}
		},
		ResolveFieldDidStartFunc: func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
			return ctx, func(interface{}, error) {}
		},
		HasResultFunc: func() bool {
			return false
		},
		GetResultFunc: func(ctx context.Context) interface{} {
			return nil
		},
	}
}

func newTestExt() *testExt {
	return &testExt{
		name: "newTextExt",
		initFn: func(ctx context.Context, p *graphql.Params) context.Context {
			return ctx
		},
		parseDidStartFn: func(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
			return ctx, func(err error) {}
		},
		validationDidStartFn: func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
			return ctx, func([]gqlerrors.FormattedError) {}
		},
		executionDidStartFn: func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
			return ctx, func(r *graphql.Result) {}
		},
		resolveFieldDidStartFn: func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
			return ctx, func(v interface{}, err error) {}
		},
		hasResultFn: func() bool {
			return false
		},
		getResultFn: func(context.Context) interface{} {
			return nil
		},
	}
}

type testExt struct {
	name                   string
	initFn                 func(ctx context.Context, p *graphql.Params) context.Context
	hasResultFn            func() bool
	getResultFn            func(context.Context) interface{}
	parseDidStartFn        func(ctx context.Context) (context.Context, graphql.ParseFinishFunc)
	validationDidStartFn   func(ctx context.Context) (context.Context, graphql.ValidationFinishFunc)
	executionDidStartFn    func(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc)
	resolveFieldDidStartFn func(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc)
}

func (t *testExt) Init(ctx context.Context, p *graphql.Params) context.Context {
	return t.initFn(ctx, p)
}

func (t *testExt) Name() string {
	return t.name
}

func (t *testExt) HasResult() bool {
	return t.hasResultFn()
}

func (t *testExt) GetResult(ctx context.Context) interface{} {
	return t.getResultFn(ctx)
}

func (t *testExt) ParseDidStart(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
	return t.parseDidStartFn(ctx)
}

func (t *testExt) ValidationDidStart(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
	return t.validationDidStartFn(ctx)
}

func (t *testExt) ExecutionDidStart(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
	return t.executionDidStartFn(ctx)
}

func (t *testExt) ResolveFieldDidStart(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
	return t.resolveFieldDidStartFn(ctx, i)
}
