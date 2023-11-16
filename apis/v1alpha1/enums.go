// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type DataReplicationErrorString string

const (
	DataReplicationErrorString_AGENT_NOT_SEEN                                DataReplicationErrorString = "AGENT_NOT_SEEN"
	DataReplicationErrorString_SNAPSHOTS_FAILURE                             DataReplicationErrorString = "SNAPSHOTS_FAILURE"
	DataReplicationErrorString_NOT_CONVERGING                                DataReplicationErrorString = "NOT_CONVERGING"
	DataReplicationErrorString_UNSTABLE_NETWORK                              DataReplicationErrorString = "UNSTABLE_NETWORK"
	DataReplicationErrorString_FAILED_TO_CREATE_SECURITY_GROUP               DataReplicationErrorString = "FAILED_TO_CREATE_SECURITY_GROUP"
	DataReplicationErrorString_FAILED_TO_LAUNCH_REPLICATION_SERVER           DataReplicationErrorString = "FAILED_TO_LAUNCH_REPLICATION_SERVER"
	DataReplicationErrorString_FAILED_TO_BOOT_REPLICATION_SERVER             DataReplicationErrorString = "FAILED_TO_BOOT_REPLICATION_SERVER"
	DataReplicationErrorString_FAILED_TO_AUTHENTICATE_WITH_SERVICE           DataReplicationErrorString = "FAILED_TO_AUTHENTICATE_WITH_SERVICE"
	DataReplicationErrorString_FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE       DataReplicationErrorString = "FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE"
	DataReplicationErrorString_FAILED_TO_CREATE_STAGING_DISKS                DataReplicationErrorString = "FAILED_TO_CREATE_STAGING_DISKS"
	DataReplicationErrorString_FAILED_TO_ATTACH_STAGING_DISKS                DataReplicationErrorString = "FAILED_TO_ATTACH_STAGING_DISKS"
	DataReplicationErrorString_FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT  DataReplicationErrorString = "FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT"
	DataReplicationErrorString_FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER DataReplicationErrorString = "FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER"
	DataReplicationErrorString_FAILED_TO_START_DATA_TRANSFER                 DataReplicationErrorString = "FAILED_TO_START_DATA_TRANSFER"
)

type DataReplicationInitiationStepName string

const (
	DataReplicationInitiationStepName_WAIT                                DataReplicationInitiationStepName = "WAIT"
	DataReplicationInitiationStepName_CREATE_SECURITY_GROUP               DataReplicationInitiationStepName = "CREATE_SECURITY_GROUP"
	DataReplicationInitiationStepName_LAUNCH_REPLICATION_SERVER           DataReplicationInitiationStepName = "LAUNCH_REPLICATION_SERVER"
	DataReplicationInitiationStepName_BOOT_REPLICATION_SERVER             DataReplicationInitiationStepName = "BOOT_REPLICATION_SERVER"
	DataReplicationInitiationStepName_AUTHENTICATE_WITH_SERVICE           DataReplicationInitiationStepName = "AUTHENTICATE_WITH_SERVICE"
	DataReplicationInitiationStepName_DOWNLOAD_REPLICATION_SOFTWARE       DataReplicationInitiationStepName = "DOWNLOAD_REPLICATION_SOFTWARE"
	DataReplicationInitiationStepName_CREATE_STAGING_DISKS                DataReplicationInitiationStepName = "CREATE_STAGING_DISKS"
	DataReplicationInitiationStepName_ATTACH_STAGING_DISKS                DataReplicationInitiationStepName = "ATTACH_STAGING_DISKS"
	DataReplicationInitiationStepName_PAIR_REPLICATION_SERVER_WITH_AGENT  DataReplicationInitiationStepName = "PAIR_REPLICATION_SERVER_WITH_AGENT"
	DataReplicationInitiationStepName_CONNECT_AGENT_TO_REPLICATION_SERVER DataReplicationInitiationStepName = "CONNECT_AGENT_TO_REPLICATION_SERVER"
	DataReplicationInitiationStepName_START_DATA_TRANSFER                 DataReplicationInitiationStepName = "START_DATA_TRANSFER"
)

type DataReplicationInitiationStepStatus string

const (
	DataReplicationInitiationStepStatus_NOT_STARTED DataReplicationInitiationStepStatus = "NOT_STARTED"
	DataReplicationInitiationStepStatus_IN_PROGRESS DataReplicationInitiationStepStatus = "IN_PROGRESS"
	DataReplicationInitiationStepStatus_SUCCEEDED   DataReplicationInitiationStepStatus = "SUCCEEDED"
	DataReplicationInitiationStepStatus_FAILED      DataReplicationInitiationStepStatus = "FAILED"
	DataReplicationInitiationStepStatus_SKIPPED     DataReplicationInitiationStepStatus = "SKIPPED"
)

type DataReplicationState string

const (
	DataReplicationState_STOPPED           DataReplicationState = "STOPPED"
	DataReplicationState_INITIATING        DataReplicationState = "INITIATING"
	DataReplicationState_INITIAL_SYNC      DataReplicationState = "INITIAL_SYNC"
	DataReplicationState_BACKLOG           DataReplicationState = "BACKLOG"
	DataReplicationState_CREATING_SNAPSHOT DataReplicationState = "CREATING_SNAPSHOT"
	DataReplicationState_CONTINUOUS        DataReplicationState = "CONTINUOUS"
	DataReplicationState_PAUSED            DataReplicationState = "PAUSED"
	DataReplicationState_RESCAN            DataReplicationState = "RESCAN"
	DataReplicationState_STALLED           DataReplicationState = "STALLED"
	DataReplicationState_DISCONNECTED      DataReplicationState = "DISCONNECTED"
)

type EC2InstanceState string

const (
	EC2InstanceState_PENDING       EC2InstanceState = "PENDING"
	EC2InstanceState_RUNNING       EC2InstanceState = "RUNNING"
	EC2InstanceState_STOPPING      EC2InstanceState = "STOPPING"
	EC2InstanceState_STOPPED       EC2InstanceState = "STOPPED"
	EC2InstanceState_SHUTTING_DOWN EC2InstanceState = "SHUTTING-DOWN"
	EC2InstanceState_TERMINATED    EC2InstanceState = "TERMINATED"
	EC2InstanceState_NOT_FOUND     EC2InstanceState = "NOT_FOUND"
)

type ExtensionStatus string

const (
	ExtensionStatus_EXTENDED        ExtensionStatus = "EXTENDED"
	ExtensionStatus_EXTENSION_ERROR ExtensionStatus = "EXTENSION_ERROR"
	ExtensionStatus_NOT_EXTENDED    ExtensionStatus = "NOT_EXTENDED"
)

type FailbackLaunchType string

const (
	FailbackLaunchType_RECOVERY FailbackLaunchType = "RECOVERY"
	FailbackLaunchType_DRILL    FailbackLaunchType = "DRILL"
)

type FailbackReplicationError string

const (
	FailbackReplicationError_AGENT_NOT_SEEN                                              FailbackReplicationError = "AGENT_NOT_SEEN"
	FailbackReplicationError_FAILBACK_CLIENT_NOT_SEEN                                    FailbackReplicationError = "FAILBACK_CLIENT_NOT_SEEN"
	FailbackReplicationError_NOT_CONVERGING                                              FailbackReplicationError = "NOT_CONVERGING"
	FailbackReplicationError_UNSTABLE_NETWORK                                            FailbackReplicationError = "UNSTABLE_NETWORK"
	FailbackReplicationError_FAILED_TO_ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION         FailbackReplicationError = "FAILED_TO_ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION"
	FailbackReplicationError_FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT  FailbackReplicationError = "FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT"
	FailbackReplicationError_FAILED_TO_CONFIGURE_REPLICATION_SOFTWARE                    FailbackReplicationError = "FAILED_TO_CONFIGURE_REPLICATION_SOFTWARE"
	FailbackReplicationError_FAILED_TO_PAIR_AGENT_WITH_REPLICATION_SOFTWARE              FailbackReplicationError = "FAILED_TO_PAIR_AGENT_WITH_REPLICATION_SOFTWARE"
	FailbackReplicationError_FAILED_TO_ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION FailbackReplicationError = "FAILED_TO_ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION"
	FailbackReplicationError_FAILED_GETTING_REPLICATION_STATE                            FailbackReplicationError = "FAILED_GETTING_REPLICATION_STATE"
	FailbackReplicationError_SNAPSHOTS_FAILURE                                           FailbackReplicationError = "SNAPSHOTS_FAILURE"
	FailbackReplicationError_FAILED_TO_CREATE_SECURITY_GROUP                             FailbackReplicationError = "FAILED_TO_CREATE_SECURITY_GROUP"
	FailbackReplicationError_FAILED_TO_LAUNCH_REPLICATION_SERVER                         FailbackReplicationError = "FAILED_TO_LAUNCH_REPLICATION_SERVER"
	FailbackReplicationError_FAILED_TO_BOOT_REPLICATION_SERVER                           FailbackReplicationError = "FAILED_TO_BOOT_REPLICATION_SERVER"
	FailbackReplicationError_FAILED_TO_AUTHENTICATE_WITH_SERVICE                         FailbackReplicationError = "FAILED_TO_AUTHENTICATE_WITH_SERVICE"
	FailbackReplicationError_FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE                     FailbackReplicationError = "FAILED_TO_DOWNLOAD_REPLICATION_SOFTWARE"
	FailbackReplicationError_FAILED_TO_CREATE_STAGING_DISKS                              FailbackReplicationError = "FAILED_TO_CREATE_STAGING_DISKS"
	FailbackReplicationError_FAILED_TO_ATTACH_STAGING_DISKS                              FailbackReplicationError = "FAILED_TO_ATTACH_STAGING_DISKS"
	FailbackReplicationError_FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT                FailbackReplicationError = "FAILED_TO_PAIR_REPLICATION_SERVER_WITH_AGENT"
	FailbackReplicationError_FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER               FailbackReplicationError = "FAILED_TO_CONNECT_AGENT_TO_REPLICATION_SERVER"
	FailbackReplicationError_FAILED_TO_START_DATA_TRANSFER                               FailbackReplicationError = "FAILED_TO_START_DATA_TRANSFER"
)

type FailbackState string

const (
	FailbackState_FAILBACK_NOT_STARTED                FailbackState = "FAILBACK_NOT_STARTED"
	FailbackState_FAILBACK_IN_PROGRESS                FailbackState = "FAILBACK_IN_PROGRESS"
	FailbackState_FAILBACK_READY_FOR_LAUNCH           FailbackState = "FAILBACK_READY_FOR_LAUNCH"
	FailbackState_FAILBACK_COMPLETED                  FailbackState = "FAILBACK_COMPLETED"
	FailbackState_FAILBACK_ERROR                      FailbackState = "FAILBACK_ERROR"
	FailbackState_FAILBACK_NOT_READY_FOR_LAUNCH       FailbackState = "FAILBACK_NOT_READY_FOR_LAUNCH"
	FailbackState_FAILBACK_LAUNCH_STATE_NOT_AVAILABLE FailbackState = "FAILBACK_LAUNCH_STATE_NOT_AVAILABLE"
)

type InitiatedBy string

const (
	InitiatedBy_START_RECOVERY               InitiatedBy = "START_RECOVERY"
	InitiatedBy_START_DRILL                  InitiatedBy = "START_DRILL"
	InitiatedBy_FAILBACK                     InitiatedBy = "FAILBACK"
	InitiatedBy_DIAGNOSTIC                   InitiatedBy = "DIAGNOSTIC"
	InitiatedBy_TERMINATE_RECOVERY_INSTANCES InitiatedBy = "TERMINATE_RECOVERY_INSTANCES"
	InitiatedBy_TARGET_ACCOUNT               InitiatedBy = "TARGET_ACCOUNT"
	InitiatedBy_CREATE_NETWORK_RECOVERY      InitiatedBy = "CREATE_NETWORK_RECOVERY"
	InitiatedBy_UPDATE_NETWORK_RECOVERY      InitiatedBy = "UPDATE_NETWORK_RECOVERY"
	InitiatedBy_ASSOCIATE_NETWORK_RECOVERY   InitiatedBy = "ASSOCIATE_NETWORK_RECOVERY"
)

type JobLogEvent string

const (
	JobLogEvent_JOB_START                           JobLogEvent = "JOB_START"
	JobLogEvent_SERVER_SKIPPED                      JobLogEvent = "SERVER_SKIPPED"
	JobLogEvent_CLEANUP_START                       JobLogEvent = "CLEANUP_START"
	JobLogEvent_CLEANUP_END                         JobLogEvent = "CLEANUP_END"
	JobLogEvent_CLEANUP_FAIL                        JobLogEvent = "CLEANUP_FAIL"
	JobLogEvent_SNAPSHOT_START                      JobLogEvent = "SNAPSHOT_START"
	JobLogEvent_SNAPSHOT_END                        JobLogEvent = "SNAPSHOT_END"
	JobLogEvent_SNAPSHOT_FAIL                       JobLogEvent = "SNAPSHOT_FAIL"
	JobLogEvent_USING_PREVIOUS_SNAPSHOT             JobLogEvent = "USING_PREVIOUS_SNAPSHOT"
	JobLogEvent_USING_PREVIOUS_SNAPSHOT_FAILED      JobLogEvent = "USING_PREVIOUS_SNAPSHOT_FAILED"
	JobLogEvent_CONVERSION_START                    JobLogEvent = "CONVERSION_START"
	JobLogEvent_CONVERSION_END                      JobLogEvent = "CONVERSION_END"
	JobLogEvent_CONVERSION_FAIL                     JobLogEvent = "CONVERSION_FAIL"
	JobLogEvent_LAUNCH_START                        JobLogEvent = "LAUNCH_START"
	JobLogEvent_LAUNCH_FAILED                       JobLogEvent = "LAUNCH_FAILED"
	JobLogEvent_JOB_CANCEL                          JobLogEvent = "JOB_CANCEL"
	JobLogEvent_JOB_END                             JobLogEvent = "JOB_END"
	JobLogEvent_DEPLOY_NETWORK_CONFIGURATION_START  JobLogEvent = "DEPLOY_NETWORK_CONFIGURATION_START"
	JobLogEvent_DEPLOY_NETWORK_CONFIGURATION_END    JobLogEvent = "DEPLOY_NETWORK_CONFIGURATION_END"
	JobLogEvent_DEPLOY_NETWORK_CONFIGURATION_FAILED JobLogEvent = "DEPLOY_NETWORK_CONFIGURATION_FAILED"
	JobLogEvent_UPDATE_NETWORK_CONFIGURATION_START  JobLogEvent = "UPDATE_NETWORK_CONFIGURATION_START"
	JobLogEvent_UPDATE_NETWORK_CONFIGURATION_END    JobLogEvent = "UPDATE_NETWORK_CONFIGURATION_END"
	JobLogEvent_UPDATE_NETWORK_CONFIGURATION_FAILED JobLogEvent = "UPDATE_NETWORK_CONFIGURATION_FAILED"
	JobLogEvent_UPDATE_LAUNCH_TEMPLATE_START        JobLogEvent = "UPDATE_LAUNCH_TEMPLATE_START"
	JobLogEvent_UPDATE_LAUNCH_TEMPLATE_END          JobLogEvent = "UPDATE_LAUNCH_TEMPLATE_END"
	JobLogEvent_UPDATE_LAUNCH_TEMPLATE_FAILED       JobLogEvent = "UPDATE_LAUNCH_TEMPLATE_FAILED"
	JobLogEvent_NETWORK_RECOVERY_FAIL               JobLogEvent = "NETWORK_RECOVERY_FAIL"
)

type JobStatus string

const (
	JobStatus_PENDING   JobStatus = "PENDING"
	JobStatus_STARTED   JobStatus = "STARTED"
	JobStatus_COMPLETED JobStatus = "COMPLETED"
)

type JobType string

const (
	JobType_LAUNCH                    JobType = "LAUNCH"
	JobType_TERMINATE                 JobType = "TERMINATE"
	JobType_CREATE_CONVERTED_SNAPSHOT JobType = "CREATE_CONVERTED_SNAPSHOT"
)

type LastLaunchResult string

const (
	LastLaunchResult_NOT_STARTED LastLaunchResult = "NOT_STARTED"
	LastLaunchResult_PENDING     LastLaunchResult = "PENDING"
	LastLaunchResult_SUCCEEDED   LastLaunchResult = "SUCCEEDED"
	LastLaunchResult_FAILED      LastLaunchResult = "FAILED"
)

type LastLaunchType string

const (
	LastLaunchType_RECOVERY LastLaunchType = "RECOVERY"
	LastLaunchType_DRILL    LastLaunchType = "DRILL"
)

type LaunchActionCategory string

const (
	LaunchActionCategory_MONITORING    LaunchActionCategory = "MONITORING"
	LaunchActionCategory_VALIDATION    LaunchActionCategory = "VALIDATION"
	LaunchActionCategory_CONFIGURATION LaunchActionCategory = "CONFIGURATION"
	LaunchActionCategory_SECURITY      LaunchActionCategory = "SECURITY"
	LaunchActionCategory_OTHER         LaunchActionCategory = "OTHER"
)

type LaunchActionParameterType string

const (
	LaunchActionParameterType_SSM_STORE LaunchActionParameterType = "SSM_STORE"
	LaunchActionParameterType_DYNAMIC   LaunchActionParameterType = "DYNAMIC"
)

type LaunchActionRunStatus string

const (
	LaunchActionRunStatus_IN_PROGRESS LaunchActionRunStatus = "IN_PROGRESS"
	LaunchActionRunStatus_SUCCEEDED   LaunchActionRunStatus = "SUCCEEDED"
	LaunchActionRunStatus_FAILED      LaunchActionRunStatus = "FAILED"
)

type LaunchActionType string

const (
	LaunchActionType_SSM_AUTOMATION LaunchActionType = "SSM_AUTOMATION"
	LaunchActionType_SSM_COMMAND    LaunchActionType = "SSM_COMMAND"
)

type LaunchDisposition string

const (
	LaunchDisposition_STOPPED LaunchDisposition = "STOPPED"
	LaunchDisposition_STARTED LaunchDisposition = "STARTED"
)

type LaunchStatus string

const (
	LaunchStatus_PENDING     LaunchStatus = "PENDING"
	LaunchStatus_IN_PROGRESS LaunchStatus = "IN_PROGRESS"
	LaunchStatus_LAUNCHED    LaunchStatus = "LAUNCHED"
	LaunchStatus_FAILED      LaunchStatus = "FAILED"
	LaunchStatus_TERMINATED  LaunchStatus = "TERMINATED"
)

type OriginEnvironment string

const (
	OriginEnvironment_ON_PREMISES OriginEnvironment = "ON_PREMISES"
	OriginEnvironment_AWS         OriginEnvironment = "AWS"
)

type PITPolicyRuleUnits string

const (
	PITPolicyRuleUnits_MINUTE PITPolicyRuleUnits = "MINUTE"
	PITPolicyRuleUnits_HOUR   PITPolicyRuleUnits = "HOUR"
	PITPolicyRuleUnits_DAY    PITPolicyRuleUnits = "DAY"
)

type RecoveryInstanceDataReplicationInitiationStepName string

const (
	RecoveryInstanceDataReplicationInitiationStepName_LINK_FAILBACK_CLIENT_WITH_RECOVERY_INSTANCE       RecoveryInstanceDataReplicationInitiationStepName = "LINK_FAILBACK_CLIENT_WITH_RECOVERY_INSTANCE"
	RecoveryInstanceDataReplicationInitiationStepName_COMPLETE_VOLUME_MAPPING                           RecoveryInstanceDataReplicationInitiationStepName = "COMPLETE_VOLUME_MAPPING"
	RecoveryInstanceDataReplicationInitiationStepName_ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION         RecoveryInstanceDataReplicationInitiationStepName = "ESTABLISH_RECOVERY_INSTANCE_COMMUNICATION"
	RecoveryInstanceDataReplicationInitiationStepName_DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT  RecoveryInstanceDataReplicationInitiationStepName = "DOWNLOAD_REPLICATION_SOFTWARE_TO_FAILBACK_CLIENT"
	RecoveryInstanceDataReplicationInitiationStepName_CONFIGURE_REPLICATION_SOFTWARE                    RecoveryInstanceDataReplicationInitiationStepName = "CONFIGURE_REPLICATION_SOFTWARE"
	RecoveryInstanceDataReplicationInitiationStepName_PAIR_AGENT_WITH_REPLICATION_SOFTWARE              RecoveryInstanceDataReplicationInitiationStepName = "PAIR_AGENT_WITH_REPLICATION_SOFTWARE"
	RecoveryInstanceDataReplicationInitiationStepName_ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION RecoveryInstanceDataReplicationInitiationStepName = "ESTABLISH_AGENT_REPLICATOR_SOFTWARE_COMMUNICATION"
	RecoveryInstanceDataReplicationInitiationStepName_WAIT                                              RecoveryInstanceDataReplicationInitiationStepName = "WAIT"
	RecoveryInstanceDataReplicationInitiationStepName_CREATE_SECURITY_GROUP                             RecoveryInstanceDataReplicationInitiationStepName = "CREATE_SECURITY_GROUP"
	RecoveryInstanceDataReplicationInitiationStepName_LAUNCH_REPLICATION_SERVER                         RecoveryInstanceDataReplicationInitiationStepName = "LAUNCH_REPLICATION_SERVER"
	RecoveryInstanceDataReplicationInitiationStepName_BOOT_REPLICATION_SERVER                           RecoveryInstanceDataReplicationInitiationStepName = "BOOT_REPLICATION_SERVER"
	RecoveryInstanceDataReplicationInitiationStepName_AUTHENTICATE_WITH_SERVICE                         RecoveryInstanceDataReplicationInitiationStepName = "AUTHENTICATE_WITH_SERVICE"
	RecoveryInstanceDataReplicationInitiationStepName_DOWNLOAD_REPLICATION_SOFTWARE                     RecoveryInstanceDataReplicationInitiationStepName = "DOWNLOAD_REPLICATION_SOFTWARE"
	RecoveryInstanceDataReplicationInitiationStepName_CREATE_STAGING_DISKS                              RecoveryInstanceDataReplicationInitiationStepName = "CREATE_STAGING_DISKS"
	RecoveryInstanceDataReplicationInitiationStepName_ATTACH_STAGING_DISKS                              RecoveryInstanceDataReplicationInitiationStepName = "ATTACH_STAGING_DISKS"
	RecoveryInstanceDataReplicationInitiationStepName_PAIR_REPLICATION_SERVER_WITH_AGENT                RecoveryInstanceDataReplicationInitiationStepName = "PAIR_REPLICATION_SERVER_WITH_AGENT"
	RecoveryInstanceDataReplicationInitiationStepName_CONNECT_AGENT_TO_REPLICATION_SERVER               RecoveryInstanceDataReplicationInitiationStepName = "CONNECT_AGENT_TO_REPLICATION_SERVER"
	RecoveryInstanceDataReplicationInitiationStepName_START_DATA_TRANSFER                               RecoveryInstanceDataReplicationInitiationStepName = "START_DATA_TRANSFER"
)

type RecoveryInstanceDataReplicationInitiationStepStatus string

const (
	RecoveryInstanceDataReplicationInitiationStepStatus_NOT_STARTED RecoveryInstanceDataReplicationInitiationStepStatus = "NOT_STARTED"
	RecoveryInstanceDataReplicationInitiationStepStatus_IN_PROGRESS RecoveryInstanceDataReplicationInitiationStepStatus = "IN_PROGRESS"
	RecoveryInstanceDataReplicationInitiationStepStatus_SUCCEEDED   RecoveryInstanceDataReplicationInitiationStepStatus = "SUCCEEDED"
	RecoveryInstanceDataReplicationInitiationStepStatus_FAILED      RecoveryInstanceDataReplicationInitiationStepStatus = "FAILED"
	RecoveryInstanceDataReplicationInitiationStepStatus_SKIPPED     RecoveryInstanceDataReplicationInitiationStepStatus = "SKIPPED"
)

type RecoveryInstanceDataReplicationState string

const (
	RecoveryInstanceDataReplicationState_STOPPED                         RecoveryInstanceDataReplicationState = "STOPPED"
	RecoveryInstanceDataReplicationState_INITIATING                      RecoveryInstanceDataReplicationState = "INITIATING"
	RecoveryInstanceDataReplicationState_INITIAL_SYNC                    RecoveryInstanceDataReplicationState = "INITIAL_SYNC"
	RecoveryInstanceDataReplicationState_BACKLOG                         RecoveryInstanceDataReplicationState = "BACKLOG"
	RecoveryInstanceDataReplicationState_CREATING_SNAPSHOT               RecoveryInstanceDataReplicationState = "CREATING_SNAPSHOT"
	RecoveryInstanceDataReplicationState_CONTINUOUS                      RecoveryInstanceDataReplicationState = "CONTINUOUS"
	RecoveryInstanceDataReplicationState_PAUSED                          RecoveryInstanceDataReplicationState = "PAUSED"
	RecoveryInstanceDataReplicationState_RESCAN                          RecoveryInstanceDataReplicationState = "RESCAN"
	RecoveryInstanceDataReplicationState_STALLED                         RecoveryInstanceDataReplicationState = "STALLED"
	RecoveryInstanceDataReplicationState_DISCONNECTED                    RecoveryInstanceDataReplicationState = "DISCONNECTED"
	RecoveryInstanceDataReplicationState_REPLICATION_STATE_NOT_AVAILABLE RecoveryInstanceDataReplicationState = "REPLICATION_STATE_NOT_AVAILABLE"
	RecoveryInstanceDataReplicationState_NOT_STARTED                     RecoveryInstanceDataReplicationState = "NOT_STARTED"
)

type RecoveryResult string

const (
	RecoveryResult_NOT_STARTED       RecoveryResult = "NOT_STARTED"
	RecoveryResult_IN_PROGRESS       RecoveryResult = "IN_PROGRESS"
	RecoveryResult_SUCCESS           RecoveryResult = "SUCCESS"
	RecoveryResult_FAIL              RecoveryResult = "FAIL"
	RecoveryResult_PARTIAL_SUCCESS   RecoveryResult = "PARTIAL_SUCCESS"
	RecoveryResult_ASSOCIATE_SUCCESS RecoveryResult = "ASSOCIATE_SUCCESS"
	RecoveryResult_ASSOCIATE_FAIL    RecoveryResult = "ASSOCIATE_FAIL"
)

type RecoverySnapshotsOrder string

const (
	RecoverySnapshotsOrder_ASC  RecoverySnapshotsOrder = "ASC"
	RecoverySnapshotsOrder_DESC RecoverySnapshotsOrder = "DESC"
)

type ReplicationConfigurationDataPlaneRouting string

const (
	ReplicationConfigurationDataPlaneRouting_PRIVATE_IP ReplicationConfigurationDataPlaneRouting = "PRIVATE_IP"
	ReplicationConfigurationDataPlaneRouting_PUBLIC_IP  ReplicationConfigurationDataPlaneRouting = "PUBLIC_IP"
)

type ReplicationConfigurationDefaultLargeStagingDiskType string

const (
	ReplicationConfigurationDefaultLargeStagingDiskType_GP2  ReplicationConfigurationDefaultLargeStagingDiskType = "GP2"
	ReplicationConfigurationDefaultLargeStagingDiskType_GP3  ReplicationConfigurationDefaultLargeStagingDiskType = "GP3"
	ReplicationConfigurationDefaultLargeStagingDiskType_ST1  ReplicationConfigurationDefaultLargeStagingDiskType = "ST1"
	ReplicationConfigurationDefaultLargeStagingDiskType_AUTO ReplicationConfigurationDefaultLargeStagingDiskType = "AUTO"
)

type ReplicationConfigurationEBSEncryption string

const (
	ReplicationConfigurationEBSEncryption_DEFAULT ReplicationConfigurationEBSEncryption = "DEFAULT"
	ReplicationConfigurationEBSEncryption_CUSTOM  ReplicationConfigurationEBSEncryption = "CUSTOM"
	ReplicationConfigurationEBSEncryption_NONE    ReplicationConfigurationEBSEncryption = "NONE"
)

type ReplicationConfigurationReplicatedDiskStagingDiskType string

const (
	ReplicationConfigurationReplicatedDiskStagingDiskType_AUTO     ReplicationConfigurationReplicatedDiskStagingDiskType = "AUTO"
	ReplicationConfigurationReplicatedDiskStagingDiskType_GP2      ReplicationConfigurationReplicatedDiskStagingDiskType = "GP2"
	ReplicationConfigurationReplicatedDiskStagingDiskType_GP3      ReplicationConfigurationReplicatedDiskStagingDiskType = "GP3"
	ReplicationConfigurationReplicatedDiskStagingDiskType_IO1      ReplicationConfigurationReplicatedDiskStagingDiskType = "IO1"
	ReplicationConfigurationReplicatedDiskStagingDiskType_SC1      ReplicationConfigurationReplicatedDiskStagingDiskType = "SC1"
	ReplicationConfigurationReplicatedDiskStagingDiskType_ST1      ReplicationConfigurationReplicatedDiskStagingDiskType = "ST1"
	ReplicationConfigurationReplicatedDiskStagingDiskType_STANDARD ReplicationConfigurationReplicatedDiskStagingDiskType = "STANDARD"
)

type ReplicationDirection string

const (
	ReplicationDirection_FAILOVER ReplicationDirection = "FAILOVER"
	ReplicationDirection_FAILBACK ReplicationDirection = "FAILBACK"
)

type ReplicationStatus string

const (
	ReplicationStatus_STOPPED     ReplicationStatus = "STOPPED"
	ReplicationStatus_IN_PROGRESS ReplicationStatus = "IN_PROGRESS"
	ReplicationStatus_PROTECTED   ReplicationStatus = "PROTECTED"
	ReplicationStatus_ERROR       ReplicationStatus = "ERROR"
)

type TargetInstanceTypeRightSizingMethod string

const (
	TargetInstanceTypeRightSizingMethod_NONE   TargetInstanceTypeRightSizingMethod = "NONE"
	TargetInstanceTypeRightSizingMethod_BASIC  TargetInstanceTypeRightSizingMethod = "BASIC"
	TargetInstanceTypeRightSizingMethod_IN_AWS TargetInstanceTypeRightSizingMethod = "IN_AWS"
)

type ValidationExceptionReason string

const (
	ValidationExceptionReason_unknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReason_cannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReason_fieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReason_other                 ValidationExceptionReason = "other"
)
