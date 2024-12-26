<br>

<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/Permify/kivo/raw/master/assets/images/logo-iam-copilot-dark.svg">
    <img alt="Kivo logo" src="https://github.com/Permify/kivo/raw/master/assets/images/logo-iam-copilot-light.svg" width="40%">
  </picture>
<h1 align="center">
   Kivo - Open Source IAM Copilot
</h1>
</div>
<p align="center">
   Kivo is an open-source CLI for managing IAM (Identity and Access Management) operations with AI assistance. It simplifies tasks like managing users, roles, and permissions and ensuring secure, consistent processes across environments.
</p>

### Installation Steps

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
