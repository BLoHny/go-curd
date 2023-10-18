// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/blohny/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/blohny/ent/tourproduct"
	"github.com/blohny/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// TourProduct is the client for interacting with the TourProduct builders.
	TourProduct *TourProductClient
	// USER is the client for interacting with the USER builders.
	USER *USERClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.TourProduct = NewTourProductClient(c.config)
	c.USER = NewUSERClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		TourProduct: NewTourProductClient(cfg),
		USER:        NewUSERClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		TourProduct: NewTourProductClient(cfg),
		USER:        NewUSERClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		TourProduct.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.TourProduct.Use(hooks...)
	c.USER.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.TourProduct.Intercept(interceptors...)
	c.USER.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *TourProductMutation:
		return c.TourProduct.mutate(ctx, m)
	case *USERMutation:
		return c.USER.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// TourProductClient is a client for the TourProduct schema.
type TourProductClient struct {
	config
}

// NewTourProductClient returns a client for the TourProduct from the given config.
func NewTourProductClient(c config) *TourProductClient {
	return &TourProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tourproduct.Hooks(f(g(h())))`.
func (c *TourProductClient) Use(hooks ...Hook) {
	c.hooks.TourProduct = append(c.hooks.TourProduct, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tourproduct.Intercept(f(g(h())))`.
func (c *TourProductClient) Intercept(interceptors ...Interceptor) {
	c.inters.TourProduct = append(c.inters.TourProduct, interceptors...)
}

// Create returns a builder for creating a TourProduct entity.
func (c *TourProductClient) Create() *TourProductCreate {
	mutation := newTourProductMutation(c.config, OpCreate)
	return &TourProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TourProduct entities.
func (c *TourProductClient) CreateBulk(builders ...*TourProductCreate) *TourProductCreateBulk {
	return &TourProductCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TourProductClient) MapCreateBulk(slice any, setFunc func(*TourProductCreate, int)) *TourProductCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TourProductCreateBulk{err: fmt.Errorf("calling to TourProductClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TourProductCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TourProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TourProduct.
func (c *TourProductClient) Update() *TourProductUpdate {
	mutation := newTourProductMutation(c.config, OpUpdate)
	return &TourProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TourProductClient) UpdateOne(tp *TourProduct) *TourProductUpdateOne {
	mutation := newTourProductMutation(c.config, OpUpdateOne, withTourProduct(tp))
	return &TourProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TourProductClient) UpdateOneID(id int) *TourProductUpdateOne {
	mutation := newTourProductMutation(c.config, OpUpdateOne, withTourProductID(id))
	return &TourProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TourProduct.
func (c *TourProductClient) Delete() *TourProductDelete {
	mutation := newTourProductMutation(c.config, OpDelete)
	return &TourProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TourProductClient) DeleteOne(tp *TourProduct) *TourProductDeleteOne {
	return c.DeleteOneID(tp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TourProductClient) DeleteOneID(id int) *TourProductDeleteOne {
	builder := c.Delete().Where(tourproduct.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TourProductDeleteOne{builder}
}

// Query returns a query builder for TourProduct.
func (c *TourProductClient) Query() *TourProductQuery {
	return &TourProductQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTourProduct},
		inters: c.Interceptors(),
	}
}

// Get returns a TourProduct entity by its id.
func (c *TourProductClient) Get(ctx context.Context, id int) (*TourProduct, error) {
	return c.Query().Where(tourproduct.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TourProductClient) GetX(ctx context.Context, id int) *TourProduct {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TourProductClient) Hooks() []Hook {
	return c.hooks.TourProduct
}

// Interceptors returns the client interceptors.
func (c *TourProductClient) Interceptors() []Interceptor {
	return c.inters.TourProduct
}

func (c *TourProductClient) mutate(ctx context.Context, m *TourProductMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TourProductCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TourProductUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TourProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TourProductDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TourProduct mutation op: %q", m.Op())
	}
}

// USERClient is a client for the USER schema.
type USERClient struct {
	config
}

// NewUSERClient returns a client for the USER from the given config.
func NewUSERClient(c config) *USERClient {
	return &USERClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *USERClient) Use(hooks ...Hook) {
	c.hooks.USER = append(c.hooks.USER, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *USERClient) Intercept(interceptors ...Interceptor) {
	c.inters.USER = append(c.inters.USER, interceptors...)
}

// Create returns a builder for creating a USER entity.
func (c *USERClient) Create() *USERCreate {
	mutation := newUSERMutation(c.config, OpCreate)
	return &USERCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of USER entities.
func (c *USERClient) CreateBulk(builders ...*USERCreate) *USERCreateBulk {
	return &USERCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *USERClient) MapCreateBulk(slice any, setFunc func(*USERCreate, int)) *USERCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &USERCreateBulk{err: fmt.Errorf("calling to USERClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*USERCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &USERCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for USER.
func (c *USERClient) Update() *USERUpdate {
	mutation := newUSERMutation(c.config, OpUpdate)
	return &USERUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *USERClient) UpdateOne(u *USER) *USERUpdateOne {
	mutation := newUSERMutation(c.config, OpUpdateOne, withUSER(u))
	return &USERUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *USERClient) UpdateOneID(id string) *USERUpdateOne {
	mutation := newUSERMutation(c.config, OpUpdateOne, withUSERID(id))
	return &USERUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for USER.
func (c *USERClient) Delete() *USERDelete {
	mutation := newUSERMutation(c.config, OpDelete)
	return &USERDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *USERClient) DeleteOne(u *USER) *USERDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *USERClient) DeleteOneID(id string) *USERDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &USERDeleteOne{builder}
}

// Query returns a query builder for USER.
func (c *USERClient) Query() *USERQuery {
	return &USERQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUSER},
		inters: c.Interceptors(),
	}
}

// Get returns a USER entity by its id.
func (c *USERClient) Get(ctx context.Context, id string) (*USER, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *USERClient) GetX(ctx context.Context, id string) *USER {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProducts queries the products edge of a USER.
func (c *USERClient) QueryProducts(u *USER) *TourProductQuery {
	query := (&TourProductClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(tourproduct.Table, tourproduct.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ProductsTable, user.ProductsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *USERClient) Hooks() []Hook {
	return c.hooks.USER
}

// Interceptors returns the client interceptors.
func (c *USERClient) Interceptors() []Interceptor {
	return c.inters.USER
}

func (c *USERClient) mutate(ctx context.Context, m *USERMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&USERCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&USERUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&USERUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&USERDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown USER mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		TourProduct, USER []ent.Hook
	}
	inters struct {
		TourProduct, USER []ent.Interceptor
	}
)