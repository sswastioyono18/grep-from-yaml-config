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

if you want to use your own config .env file, you can specify
```
--config test-config.env
```
## Screenshot
![Screen Shot 2021-09-08 at 13 15 53](https://user-images.githubusercontent.com/11158339/132456284-7e242bed-8388-4415-97c5-fa0d8f9abab3.png)


## TODO
- Complexity seems high for grep from config yaml since we are looping 
    - from list of app 
    - then list of target yaml
    - then list of yaml 
    - then grep content of go file.
- Unit testing
- Implement same functionality for .env & toml file
