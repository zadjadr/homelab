import os
import subprocess
import time

WAIT_TIME_SEC = 10

all_az = os.environ.get(
    "OCI_AZONES",
    "ldMg:EU-FRANKFURT-1-AD-1 ldMg:EU-FRANKFURT-1-AD-2 ldMg:EU-FRANKFURT-1-AD-3",
).split(" ")
all_names = os.environ.get("OCI_INSTANCE_NAMES", "3.cloud 4.cloud").split(" ")

while True:
    for az, name in zip(all_az, all_names):
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

        if result.returncode == 0:
            print("Succeeded!")
            exit(0)

        print(f"Waiting for {WAIT_TIME_SEC} secs before applying again.")
        time.sleep(WAIT_TIME_SEC)
