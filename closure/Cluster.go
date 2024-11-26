package closure

import "github.com/SwanHtetAungPhyo/closure/utils"

type Cluster struct {
	prefix     string
	router     *Router
	middleware []Middleware
}

func NewCluster(prefix string, router *Router, mw ...Middleware) *Cluster {
	return &Cluster{
		prefix:     prefix,
		router:     router,
		middleware: mw,
	}
}

func (c *Cluster) AddMiddleware(mw Middleware) *Cluster {
	c.middleware = append(c.middleware, mw)
	return c
}




func (c *Cluster) registerRoute(method, path string, handler Handler) {
	fullpath := utils.FullPath(c.prefix, path)
	ultimateHandler := handler
	
	for i := len(c.middleware) - 1; i >= 0; i-- {

		ultimateHandler = c.middleware[i].Apply(ultimateHandler)
	}

	c.router.Register(method, fullpath, ultimateHandler)
}


func (c *Cluster) Get(path string, handler Handler) *Cluster {
	c.registerRoute("GET", path, handler)
	return c
}

func (c *Cluster) Post(path string, handler Handler) *Cluster {
	c.registerRoute("POST", path, handler)
	return c
}

func (c *Cluster) Put(path string, handler Handler) *Cluster {
	c.registerRoute("PUT", path, handler)
	return c
}

func (c *Cluster) Patch(path string, handler Handler) *Cluster {
	c.registerRoute("PATCH", path, handler)
	return c
}

func (c *Cluster) Delete(path string, handler Handler) *Cluster {
	c.registerRoute("DELETE", path, handler)
	return c
}

func (c *Cluster) Head(path string, handler Handler) *Cluster {
	c.registerRoute("HEAD", path, handler)
	return c
}

func (c *Cluster) Options(path string, handler Handler) *Cluster {
	c.registerRoute("OPTIONS", path, handler)
	return c
}

func (c *Cluster) Trace(path string, handler Handler) *Cluster {
	c.registerRoute("TRACE", path, handler)
	return c
}