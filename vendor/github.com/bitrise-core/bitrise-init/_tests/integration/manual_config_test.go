package integration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bitrise-core/bitrise-init/models"
	"github.com/bitrise-core/bitrise-init/steps"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/stretchr/testify/require"
)

func TestManualConfig(t *testing.T) {
	tmpDir, err := pathutil.NormalizedOSTempDirPath("__manual-config__")
	require.NoError(t, err)

	t.Log("manual-config")
	{
		manualConfigDir := filepath.Join(tmpDir, "manual-config")
		require.NoError(t, os.MkdirAll(manualConfigDir, 0777))
		fmt.Printf("manualConfigDir: %s\n", manualConfigDir)

		cmd := command.New(binPath(), "--ci", "manual-config", "--output-dir", manualConfigDir)
		out, err := cmd.RunAndReturnTrimmedCombinedOutput()
		require.NoError(t, err, out)

		scanResultPth := filepath.Join(manualConfigDir, "result.yml")

		result, err := fileutil.ReadStringFromFile(scanResultPth)
		require.NoError(t, err)

		validateConfigExpectation(t, "manual-config", strings.TrimSpace(customConfigResultYML), strings.TrimSpace(result), customConfigVersions...)
	}
}

var customConfigVersions = []interface{}{
	// android
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.CachePullVersion,
	steps.ScriptVersion,
	steps.InstallMissingAndroidToolsVersion,
	steps.ChangeAndroidVersionCodeAndVersionNameVersion,
	steps.AndroidLintVersion,
	steps.AndroidUnitTestVersion,
	steps.AndroidBuildVersion,
	steps.SignAPKVersion,
	steps.DeployToBitriseIoVersion,
	steps.CachePushVersion,

	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.CachePullVersion,
	steps.ScriptVersion,
	steps.InstallMissingAndroidToolsVersion,
	steps.AndroidLintVersion,
	steps.AndroidUnitTestVersion,
	steps.DeployToBitriseIoVersion,
	steps.CachePushVersion,

	// cordova
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.NpmVersion,
	steps.GenerateCordovaBuildConfigVersion,
	steps.CordovaArchiveVersion,
	steps.DeployToBitriseIoVersion,

	// fastlane
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.FastlaneVersion,
	steps.DeployToBitriseIoVersion,

	// ionic
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.NpmVersion,
	steps.GenerateCordovaBuildConfigVersion,
	steps.IonicArchiveVersion,
	steps.DeployToBitriseIoVersion,

	// ios
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.CachePullVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.RecreateUserSchemesVersion,
	steps.CocoapodsInstallVersion,
	steps.XcodeTestVersion,
	steps.XcodeArchiveVersion,
	steps.DeployToBitriseIoVersion,
	steps.CachePushVersion,

	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.CachePullVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.RecreateUserSchemesVersion,
	steps.CocoapodsInstallVersion,
	steps.XcodeTestVersion,
	steps.DeployToBitriseIoVersion,
	steps.CachePushVersion,

	// macos
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.CachePullVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.RecreateUserSchemesVersion,
	steps.CocoapodsInstallVersion,
	steps.XcodeTestMacVersion,
	steps.XcodeArchiveMacVersion,
	steps.DeployToBitriseIoVersion,
	steps.CachePushVersion,

	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.CachePullVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.RecreateUserSchemesVersion,
	steps.CocoapodsInstallVersion,
	steps.XcodeTestMacVersion,
	steps.DeployToBitriseIoVersion,
	steps.CachePushVersion,

	// other
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.DeployToBitriseIoVersion,

	// react native
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.NpmVersion,
	steps.InstallMissingAndroidToolsVersion,
	steps.AndroidBuildVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.XcodeArchiveVersion,
	steps.DeployToBitriseIoVersion,

	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.NpmVersion,
	steps.NpmVersion,
	steps.DeployToBitriseIoVersion,

	// react native expo
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.NpmVersion,
	steps.NpmVersion,
	steps.InstallMissingAndroidToolsVersion,
	steps.AndroidBuildVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.XcodeArchiveVersion,
	steps.DeployToBitriseIoVersion,

	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.NpmVersion,
	steps.NpmVersion,
	steps.DeployToBitriseIoVersion,

	// xamarin
	models.FormatVersion,
	steps.ActivateSSHKeyVersion,
	steps.GitCloneVersion,
	steps.ScriptVersion,
	steps.CertificateAndProfileInstallerVersion,
	steps.XamarinUserManagementVersion,
	steps.NugetRestoreVersion,
	steps.XamarinComponentsRestoreVersion,
	steps.XamarinArchiveVersion,
	steps.DeployToBitriseIoVersion,
}

var customConfigResultYML = fmt.Sprintf(`options:
  android:
    title: The root directory of an Android project
    env_key: PROJECT_LOCATION
    value_map:
      _:
        title: Module
        env_key: MODULE
        value_map:
          _:
            title: Variant for building
            env_key: BUILD_VARIANT
            value_map:
              _:
                title: Variant for testing
                env_key: TEST_VARIANT
                value_map:
                  _:
                    config: default-android-config
  cordova:
    title: Directory of Cordova Config.xml
    env_key: CORDOVA_WORK_DIR
    value_map:
      _:
        title: Platform to use in cordova-cli commands
        env_key: CORDOVA_PLATFORM
        value_map:
          android:
            config: default-cordova-config
          ios:
            config: default-cordova-config
          ios,android:
            config: default-cordova-config
  fastlane:
    title: Working directory
    env_key: FASTLANE_WORK_DIR
    value_map:
      _:
        title: Fastlane lane
        env_key: FASTLANE_LANE
        value_map:
          _:
            config: default-fastlane-config
  ionic:
    title: Directory of Ionic Config.xml
    env_key: IONIC_WORK_DIR
    value_map:
      _:
        title: Platform to use in ionic-cli commands
        env_key: IONIC_PLATFORM
        value_map:
          android:
            config: default-ionic-config
          ios:
            config: default-ionic-config
          ios,android:
            config: default-ionic-config
  ios:
    title: Project (or Workspace) path
    env_key: BITRISE_PROJECT_PATH
    value_map:
      _:
        title: Scheme name
        env_key: BITRISE_SCHEME
        value_map:
          _:
            title: ipa export method
            env_key: BITRISE_EXPORT_METHOD
            value_map:
              ad-hoc:
                config: default-ios-config
              app-store:
                config: default-ios-config
              development:
                config: default-ios-config
              enterprise:
                config: default-ios-config
  macos:
    title: Project (or Workspace) path
    env_key: BITRISE_PROJECT_PATH
    value_map:
      _:
        title: Scheme name
        env_key: BITRISE_SCHEME
        value_map:
          _:
            title: |-
              Application export method
              NOTE: `+"`"+`none`+"`"+` means: Export a copy of the application without re-signing.
            env_key: BITRISE_EXPORT_METHOD
            value_map:
              app-store:
                config: default-macos-config
              developer-id:
                config: default-macos-config
              development:
                config: default-macos-config
              none:
                config: default-macos-config
  react-native:
    title: The root directory of an Android project
    env_key: PROJECT_LOCATION
    value_map:
      _:
        title: Module
        env_key: MODULE
        value_map:
          _:
            title: Variant for building
            env_key: BUILD_VARIANT
            value_map:
              _:
                title: Project (or Workspace) path
                env_key: BITRISE_PROJECT_PATH
                value_map:
                  _:
                    title: Scheme name
                    env_key: BITRISE_SCHEME
                    value_map:
                      _:
                        title: ipa export method
                        env_key: BITRISE_EXPORT_METHOD
                        value_map:
                          ad-hoc:
                            config: default-react-native-config
                          app-store:
                            config: default-react-native-config
                          development:
                            config: default-react-native-config
                          enterprise:
                            config: default-react-native-config
  react-native-expo:
    title: The root directory of an Android project
    env_key: PROJECT_LOCATION
    value_map:
      _:
        title: Module
        env_key: MODULE
        value_map:
          _:
            title: Variant for building
            env_key: BUILD_VARIANT
            value_map:
              _:
                title: Project (or Workspace) path
                env_key: BITRISE_PROJECT_PATH
                value_map:
                  _:
                    title: Scheme name
                    env_key: BITRISE_SCHEME
                    value_map:
                      _:
                        title: ipa export method
                        env_key: BITRISE_EXPORT_METHOD
                        value_map:
                          ad-hoc:
                            config: default-react-native-expo-config
                          app-store:
                            config: default-react-native-expo-config
                          development:
                            config: default-react-native-expo-config
                          enterprise:
                            config: default-react-native-expo-config
  xamarin:
    title: Path to the Xamarin Solution file
    env_key: BITRISE_PROJECT_PATH
    value_map:
      _:
        title: Xamarin solution configuration
        env_key: BITRISE_XAMARIN_CONFIGURATION
        value_map:
          _:
            title: Xamarin solution platform
            env_key: BITRISE_XAMARIN_PLATFORM
            value_map:
              _:
                config: default-xamarin-config
configs:
  android:
    default-android-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: android
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        deploy:
          description: |
            ## How to get a signed APK

            This workflow contains the **Sign APK** step. To sign your APK all you have to do is to:

            1. Click on **Code Signing** tab
            1. Find the **ANDROID KEYSTORE FILE** section
            1. Click or drop your file on the upload file field
            1. Fill the displayed 3 input fields:
             1. **Keystore password**
             1. **Keystore alias**
             1. **Private key password**
            1. Click on **[Save metadata]** button

            That's it! From now on, **Sign APK** step will receive your uploaded files.

            ## To run this workflow

            If you want to run this workflow manually:

            1. Open the app's build list page
            2. Click on **[Start/Schedule a Build]** button
            3. Select **deploy** in **Workflow** dropdown input
            4. Click **[Start Build]** button

            Or if you need this workflow to be started by a GIT event:

            1. Click on **Triggers** tab
            2. Setup your desired event (push/tag/pull) and select **deploy** workflow
            3. Click on **[Done]** and then **[Save]** buttons

            The next change in your repository that matches any of your trigger map event will start **deploy** workflow.
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - cache-pull@%s: {}
          - script@%s:
              title: Do anything with Script step
          - install-missing-android-tools@%s:
              inputs:
              - gradlew_path: $PROJECT_LOCATION/gradlew
          - change-android-versioncode-and-versionname@%s:
              inputs:
              - build_gradle_path: $PROJECT_LOCATION/$MODULE/build.gradle
          - android-lint@%s:
              inputs:
              - project_location: $PROJECT_LOCATION
              - module: $MODULE
              - variant: $TEST_VARIANT
          - android-unit-test@%s:
              inputs:
              - project_location: $PROJECT_LOCATION
              - module: $MODULE
              - variant: $TEST_VARIANT
          - android-build@%s:
              inputs:
              - project_location: $PROJECT_LOCATION
              - module: $MODULE
              - variant: $BUILD_VARIANT
          - sign-apk@%s:
              run_if: '{{getenv "BITRISEIO_ANDROID_KEYSTORE_URL" | ne ""}}'
          - deploy-to-bitrise-io@%s: {}
          - cache-push@%s: {}
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - cache-pull@%s: {}
          - script@%s:
              title: Do anything with Script step
          - install-missing-android-tools@%s:
              inputs:
              - gradlew_path: $PROJECT_LOCATION/gradlew
          - android-lint@%s:
              inputs:
              - project_location: $PROJECT_LOCATION
              - module: $MODULE
              - variant: $TEST_VARIANT
          - android-unit-test@%s:
              inputs:
              - project_location: $PROJECT_LOCATION
              - module: $MODULE
              - variant: $TEST_VARIANT
          - deploy-to-bitrise-io@%s: {}
          - cache-push@%s: {}
  cordova:
    default-cordova-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: cordova
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - npm@%s:
              inputs:
              - command: install
              - workdir: $CORDOVA_WORK_DIR
          - generate-cordova-build-configuration@%s: {}
          - cordova-archive@%s:
              inputs:
              - workdir: $CORDOVA_WORK_DIR
              - platform: $CORDOVA_PLATFORM
              - target: emulator
          - deploy-to-bitrise-io@%s: {}
  fastlane:
    default-fastlane-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: fastlane
      app:
        envs:
        - FASTLANE_XCODE_LIST_TIMEOUT: "120"
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - fastlane@%s:
              inputs:
              - lane: $FASTLANE_LANE
              - work_dir: $FASTLANE_WORK_DIR
          - deploy-to-bitrise-io@%s: {}
  ionic:
    default-ionic-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: ionic
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - npm@%s:
              inputs:
              - command: install
              - workdir: $IONIC_WORK_DIR
          - generate-cordova-build-configuration@%s: {}
          - ionic-archive@%s:
              inputs:
              - workdir: $IONIC_WORK_DIR
              - platform: $IONIC_PLATFORM
              - target: emulator
          - deploy-to-bitrise-io@%s: {}
  ios:
    default-ios-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: ios
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        deploy:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - cache-pull@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - recreate-user-schemes@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
          - cocoapods-install@%s: {}
          - xcode-test@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
          - xcode-archive@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
              - export_method: $BITRISE_EXPORT_METHOD
          - deploy-to-bitrise-io@%s: {}
          - cache-push@%s: {}
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - cache-pull@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - recreate-user-schemes@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
          - cocoapods-install@%s: {}
          - xcode-test@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
          - deploy-to-bitrise-io@%s: {}
          - cache-push@%s: {}
  macos:
    default-macos-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: macos
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        deploy:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - cache-pull@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - recreate-user-schemes@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
          - cocoapods-install@%s: {}
          - xcode-test-mac@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
          - xcode-archive-mac@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
              - export_method: $BITRISE_EXPORT_METHOD
          - deploy-to-bitrise-io@%s: {}
          - cache-push@%s: {}
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - cache-pull@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - recreate-user-schemes@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
          - cocoapods-install@%s: {}
          - xcode-test-mac@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
          - deploy-to-bitrise-io@%s: {}
          - cache-push@%s: {}
  other:
    other-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: other
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - deploy-to-bitrise-io@%s: {}
  react-native:
    default-react-native-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: react-native
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        deploy:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - npm@%s:
              inputs:
              - command: install
          - install-missing-android-tools@%s:
              inputs:
              - gradlew_path: $PROJECT_LOCATION/gradlew
          - android-build@%s:
              inputs:
              - project_location: $PROJECT_LOCATION
              - module: $MODULE
              - variant: $BUILD_VARIANT
          - certificate-and-profile-installer@%s: {}
          - xcode-archive@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
              - export_method: $BITRISE_EXPORT_METHOD
              - configuration: Release
          - deploy-to-bitrise-io@%s: {}
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - npm@%s:
              inputs:
              - command: install
          - npm@%s:
              inputs:
              - command: test
          - deploy-to-bitrise-io@%s: {}
  react-native-expo:
    default-react-native-expo-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: react-native-expo
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        deploy:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - npm@%s:
              inputs:
              - command: install
          - npm@%s:
              inputs:
              - command: run eject
          - install-missing-android-tools@%s:
              inputs:
              - gradlew_path: $PROJECT_LOCATION/gradlew
          - android-build@%s:
              inputs:
              - project_location: $PROJECT_LOCATION
              - module: $MODULE
              - variant: $BUILD_VARIANT
          - certificate-and-profile-installer@%s: {}
          - xcode-archive@%s:
              inputs:
              - project_path: $BITRISE_PROJECT_PATH
              - scheme: $BITRISE_SCHEME
              - export_method: $BITRISE_EXPORT_METHOD
              - configuration: Release
          - deploy-to-bitrise-io@%s: {}
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - npm@%s:
              inputs:
              - command: install
          - npm@%s:
              inputs:
              - command: test
          - deploy-to-bitrise-io@%s: {}
  xamarin:
    default-xamarin-config: |
      format_version: "%s"
      default_step_lib_source: https://github.com/bitrise-io/bitrise-steplib.git
      project_type: xamarin
      trigger_map:
      - push_branch: '*'
        workflow: primary
      - pull_request_source_branch: '*'
        workflow: primary
      workflows:
        primary:
          steps:
          - activate-ssh-key@%s:
              run_if: '{{getenv "SSH_RSA_PRIVATE_KEY" | ne ""}}'
          - git-clone@%s: {}
          - script@%s:
              title: Do anything with Script step
          - certificate-and-profile-installer@%s: {}
          - xamarin-user-management@%s:
              run_if: .IsCI
          - nuget-restore@%s: {}
          - xamarin-components-restore@%s: {}
          - xamarin-archive@%s:
              inputs:
              - xamarin_solution: $BITRISE_PROJECT_PATH
              - xamarin_configuration: $BITRISE_XAMARIN_CONFIGURATION
              - xamarin_platform: $BITRISE_XAMARIN_PLATFORM
          - deploy-to-bitrise-io@%s: {}
`, customConfigVersions...)
