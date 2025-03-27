# InMemoryCacheDemo
Cache package selection:  I went with go-cache because it because of its thread safety, expiration support, and auto clean up.  sync.Mutex and sync.Map are good choices but didn't have expiration support.

I went with using an API because it's easy to implement in memory caching and can easily run the app to test the code and do unit tests.

What I would do differently if I had to scale the app:

In-Memory Caching doesn't scale well when running multiple instances.  I would use Redis as a managed service or self-hosted.  The advantage would be that it is persistent and can be shared among multiple instances.

The app runs on a single instance.  If traffic increases it's better to containerize the app using Docker, deploy the app to kubernetes, and implement autoscaling

Since the data in a cache is usually persisted, we should implement a database to store the data.  And if the cache is expired, we can make a call to the database to get the data and store it in the cache.

Implement cache strategies for better performance(TTL, cache warming, and cache expiration handling)

Implement Middleware or a helper function to centralize error handling

Input Validation

Add Logging to handle key events like cache hits, misses, and errors

Hardcode values to a configuration file

Rate Limiting

Use Https
