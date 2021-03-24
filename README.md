# Role Policies

## Summary
EPA digital policies package

## Motivation
* To implement roles with casbin we need to specify which policies we want to implement. This package is intended to hold every policy, so changes in policies 
would only require changes in this package.

## Package structure
+ Policies 
  + policy.go
  + Single policy package -> Each policy will have its own folder
    + single_policy.conf  -> Each policy requires a .conf file
    + single_policy.csv   -> Each policy requires a .csv file
  
## API
To start using a policy from the package you just have to import the package

```golang
 "github.com/epa-datos/policies"

```
The package has the following constants:
- **TestPolicy**: Implements a policy for testing purposes

Each constant is of type policy, which has two methods:   
### GetPolicyConfPath ###
Returns the file path for the .conf file of the policy
```golang
  confFilePath := policy.GetPolicyConfPath() 
```
### GetPolicyCsvPath ###   
Returns the file path for the .csv file of the policy
```golang
  confFilePath := policy.GetPolicyCsvPath() 
```
