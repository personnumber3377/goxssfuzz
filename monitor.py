import os
import time

import json

from selenium import webdriver
from selenium.webdriver.chrome.service import Service as ChromeService
from selenium.webdriver.common.by import By
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import NoAlertPresentException


# Path to the input file that the program monitors
INPUT_FILE = "input.txt"

def get_file_modification_time(file_path):
    """Get the last modification time of a file."""
    return os.path.getmtime(file_path)

def read_input_from_file(file_path):
    """Read the input data from the given file."""
    with open(file_path, "r") as f:
        return f.read().strip()


# driver.get("data:text/html;charset=utf-8,<html><body></body></html>")



def process_input_with_selenium(driver, input_data):
    """Process the input data using Selenium."""
    #print(f"Processing input: {input_data}")

    # Example of navigating to a page and performing some action
    #driver.get("https://example.com")  # Replace with your target URL
    #print("Navigated to example.com")

    # Here, you can interact with the web page as needed
    # For example, send the input data to a search field or form
    try:
        # driver.get("data:text/html;charset=utf-8," + input_data)

        escaped_data = json.dumps(input_data)

        # Use JavaScript to replace the entire document's content
        script = f"""
        document.open();
        document.write({escaped_data});
        document.close();
        """
        driver.execute_script(script)







        # Check for an alert
        try:
            alert = driver.switch_to.alert  # Switch to alert if present
            alert_text = alert.text  # Get the alert's text
            alert.dismiss()  # Close the alert
            print(f"Alert detected! Content: {alert_text}")
            # We noticed an alert!!!
            # Just do the bullshit.
            fh = open("exploit.txt", "w")
            fh.write(input_data)
            fh.close()
            fh = open("result.txt", "w")
            fh.write("1")
            fh.close()
            #assert False
            exit(1)
            return alert_text
        except NoAlertPresentException:
            #print("No alert detected.")
            fh = open("result.txt", "w")
            fh.write("0")
            fh.close()
            return None
    except Exception as e:
        print(f"Error interacting with the web page: {e}")


def setup_driver():
    chrome_options = Options()
    chrome_options.headless = True  # Run in headless mode (no UI)
    service = Service("./chromedriver")  # Replace with your chromedriver path
    driver = webdriver.Chrome(service=service, options=chrome_options)
    return driver


SLEEP_THING = 0

def main():
    # Initialize the Selenium WebDriver once
    print("Initializing Selenium WebDriver...")
    driver = setup_driver() # webdriver.Chrome(service=ChromeService())  # Adjust if using Firefox, Edge, etc.
    print("Selenium WebDriver initialized.")
    driver.get("data:text/html;charset=utf-8,<html><body></body></html>") # Load up an initial empty page.
    # Track file modification time
    last_mod_time = 0

    # Main loop: Monitor the file for changes
    try:
        while True:


            # Just skip checking for change. This just adds performance overhead.
            '''
            if os.path.exists(INPUT_FILE):
                current_mod_time = get_file_modification_time(INPUT_FILE)

                # Check if the file was modified
                if current_mod_time > last_mod_time:
                    print("Change detected in input file.")
                    last_mod_time = current_mod_time

                    # Read and process new input
            '''

            input_data = read_input_from_file(INPUT_FILE)
            # print("poopoo")
            if input_data:
                process_input_with_selenium(driver, input_data)
            #else:
            #    print("Input file is empty.")

            # Sleep briefly before checking again
            '''
            if SLEEP_THING == 0:
                continue
            else:
                time.sleep(SLEEP_THING)
            '''

    except KeyboardInterrupt:
        print("Stopping the program.")
    finally:
        # Clean up and close the WebDriver
        driver.quit()
        print("Selenium WebDriver closed.")

if __name__ == "__main__":
    main()