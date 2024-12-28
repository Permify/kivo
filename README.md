<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/Permify/kivo/raw/master/assets/images/logo-iam-copilot-dark.png">
    <img alt="Kivo logo" src="https://github.com/Permify/kivo/raw/master/assets/images/logo-iam-copilot-light.png" width="40%">
  </picture>
<h1 align="center">
   Kivo - Open Source IAM Copilot
</h1>
</div>
<p>
Kivo is an open-source CLI for managing IAM (Identity and Access Management) operations with AI assistance.

DevOps engineers use Kivo to configure how employees in their organization access infrastructure resources. Kivo simplifies and accelerates granting and revoking access, while supporting custom policy creation — eliminating the need for tedious back-and-forth UI work.
</p>

<p align="center">
    <a href="https://github.com/Permify/kivo" target="_blank"><img src="https://img.shields.io/github/go-mod/go-version/Permify/kivo?style=for-the-badge&logo=go" alt="Permify Go Version" /></a>&nbsp;
 <a href="https://goreportcard.com/report/github.com/Permify/kivo" target="_blank"><img src="https://goreportcard.com/badge/github.com/Permify/kivo?style=for-the-badge&logo=go" alt="Kivo Go Report Card" /></a>&nbsp;
    <a href="https://github.com/Permify/kivo" target="_blank"><img src="https://img.shields.io/github/license/Permify/kivo?style=for-the-badge" alt="Kivo Licence" /></a>&nbsp;
    <a href="https://discord.gg/n6KfzYxhPp" target="_blank"><img src="https://img.shields.io/discord/950799928047833088?style=for-the-badge&logo=discord&label=DISCORD" alt="Permify Discord Channel" /></a>&nbsp;
    <a href="https://github.com/Permify/kivo/releases" target="_blank"><img src="https://img.shields.io/github/v/release/permify/kivo?include_prereleases&style=for-the-badge" alt="Kivo Release" /></a>&nbsp;
    <a href="https://img.shields.io/github/commit-activity/m/Permify/kivo?style=for-the-badge" target="_blank"><img src="https://img.shields.io/github/commit-activity/m/Permify/kivo?style=for-the-badge" alt="Kivo Commit Activity" /></a>&nbsp;
    <a href="https://img.shields.io/github/actions/workflow/status/Permify/kivo/release.yml?style=for-the-badge" target="_blank"><img src="https://img.shields.io/github/actions/workflow/status/Permify/kivo/release.yml?style=for-the-badge" alt="GitHub Workflow Status" /></a>&nbsp;
</p>  

![Kivo Demo](assets/images/kivo.gif)

## How it Works?

1. Configure your cloud credentials to enable Kivo to access resources in your infrastructure. Currently, Kivo supports AWS only.
2. Start an access flow or use AI to create an access command to fullfil access request.
3. Preview the access action and complete granting and revoking access.

### Create an Access Command with AI 

Describe the access action you want to perform. For example, "give S3 read-only access to user Omer." 

Kivo analyzes the request and generates the necessary access command using AI.

![kivo-ai-flow](https://github.com/user-attachments/assets/a64b7aab-1fe1-4cba-b5f1-bb407dab8383)

### Start an Access Flow Manually

You can also manually start any flow to complete an access action. 

There are three main flows:
   - `~ % kivo aws users`   | Grant or revoke access to/from a user.
   - `~ % kivo aws groups`  | Add or remove a user from a group.
   - `~ % kivo aws roles`   | Attach or detach a role to/from a user.

Let's repeat the example above of granting s3 read-only access to user Omer.

We will use following command to start **user** flow: `~ % kivo aws users`.

The user access flow begins by listing the users in the system. Select the user to take action on.

![select-user](https://github.com/user-attachments/assets/e764d8fd-7cce-4229-a119-45daa8ce6442)

After selecting the user, choose the operation to perform. Let’s attach a policy to user Omer.

![Screen Shot 2024-12-28 at 20 33 00](https://github.com/user-attachments/assets/71fe2f72-321c-4b64-b36f-f3574df05854)

In the next step, select the policy you want to attach. You can use "filters" in each section to search what you need.

![Screen Shot 2024-12-28 at 20 32 18](https://github.com/user-attachments/assets/e3353b23-7c43-4a66-80b9-2359c72b5934)

Finally, preview the access action.

![Screen Shot 2024-12-28 at 20 35 40](https://github.com/user-attachments/assets/6883904e-edbe-4e05-9856-c70760996adc)


## Installation Steps

1. **Install Kivo CLI:**
   ```shell
   brew tap permify/tap-kivo
   brew install kivo
   ```

2. **Set Up AWS Credentials:**

   Kivo requires AWS credentials to be configured in the file `~/.aws/credentials`. Follow these steps:

    - Create or open the `~/.aws/credentials` file using a text editor:
      ```shell
      nano ~/.aws/credentials
      ```

    - Add your AWS credentials in the following format:
      ```plaintext
      [default]
      aws_access_key_id = your_access_key
      aws_secret_access_key = your_secret_key
      ```

    - Save the file and exit (in nano, press `CTRL + O` to save, then `CTRL + X` to exit).

3. **Verify the Configuration:**

   Run the following command to confirm the credentials are set correctly:
   ```shell
   aws sts get-caller-identity
   ```
   This should return information about your AWS account. If it fails, double-check the credentials file for accuracy.

4. **Set the Default Region (Optional):**

   If your tool requires a specific AWS region, you can set it in the `~/.aws/config` file:
   ```shell
   nano ~/.aws/config
   ```
   Add:
   ```plaintext
   [default]
   region = us-east-1
   ```
   Replace `us-east-1` with your desired region.

## Communication Channels

If you like Kivo, please consider giving us a :star:

<p align="left">
<a href="https://discord.gg/n6KfzYxhPp">
 <img height="70px" width="70px" alt="permify | Discord" src="https://user-images.githubusercontent.com/39353278/187209316-3d01a799-c51b-4eaa-8f52-168047078a14.png" />
</a>
<a href="https://twitter.com/GetPermify">
  <img height="70px" width="70px" alt="permify | Twitter" src="https://user-images.githubusercontent.com/39353278/187209323-23f14261-d406-420d-80eb-1aa707a71043.png"/>
</a>
<a href="https://www.linkedin.com/company/permifyco">
  <img height="70px" width="70px" alt="permify | Linkedin" src="https://user-images.githubusercontent.com/39353278/187209321-03293a24-6f63-4321-b362-b0fc89fdd879.png" />
</a>
</p>
