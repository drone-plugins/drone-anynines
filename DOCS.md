Use this plugin for deplying an application to Anynines. You can override the
default configuration with the following parameters:

* `username` - Username for Anynines
* `password` - Password for Anynines
* `organization` - Target organization for Anynines
* `space` - Target space for Anynines

## Example

```yaml
deploy:
  anynines:
    username: foo
    password: bar
    organization: drone
    space: test
```
