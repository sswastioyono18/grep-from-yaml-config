# grep-from-yaml-config

## Why
- Project runs OK but too many unused config?
- Ever found it boring to cleanup config yaml file?

This project aim to solve that.

Grep key from config file (currently supports yaml only) and find usage in the project.

## How To
To find key not used in yaml config
```
./grep-from-yaml-config clean-yaml-config 
```

To find key not used in yaml secret
```
./grep-from-yaml-config clean-yaml-secret 
```

## TODO
Complexity seems high for grep from config yaml since we are looping from list of app, then list of target yaml, then list of yaml and then grep content of go file.

