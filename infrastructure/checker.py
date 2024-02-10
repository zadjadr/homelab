import os
import subprocess
import time
from itertools import product

WAIT_TIME_SEC = 10

# Get environment variables with default values
all_az = os.environ.get(
    "OCI_AZONES",
    "ldMg:EU-FRANKFURT-1-AD-1 ldMg:EU-FRANKFURT-1-AD-2 ldMg:EU-FRANKFURT-1-AD-3",
).split()
all_names = os.environ.get("OCI_INSTANCE_NAMES", "3.cloud 4.cloud").split()


def apply_terraform(az, name):
    """
    Apply Terraform with given availability domain and instance name.
    """
    print(f"Checking for {name} in {az}")
    result = subprocess.run(
        [
            "terraform",
            "apply",
            "-auto-approve",
            "-var",
            f"availability_domain={az}",
            "-var",
            f"instance_name={name}",
        ],
        stdout=subprocess.DEVNULL,
    )
    return result.returncode == 0


def main():
    """
    Main function to iterate over all_az and all_names and apply Terraform.
    """
    for az, name in product(all_az, all_names):
        if apply_terraform(az, name):
            print("Succeeded!")
            return
        print(f"Waiting for {WAIT_TIME_SEC} secs before applying again.")
        time.sleep(WAIT_TIME_SEC)


if __name__ == "__main__":
    main()
