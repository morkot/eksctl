package cloudformation

// AWSOpsWorksApp_EnvironmentVariable AWS CloudFormation Resource (AWS::OpsWorks::App.EnvironmentVariable)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html
type AWSOpsWorksApp_EnvironmentVariable struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-key
	Key *StringIntrinsic `json:"Key,omitempty"`

	// Secure AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-secure
	Secure bool `json:"Secure,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#value
	Value *StringIntrinsic `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksApp_EnvironmentVariable) AWSCloudFormationType() string {
	return "AWS::OpsWorks::App.EnvironmentVariable"
}
