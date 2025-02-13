from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import NoAlertPresentException

# Configure Selenium with a headless browser
def setup_driver():
    chrome_options = Options()
    chrome_options.headless = True  # Run in headless mode (no UI)
    service = Service("./chromedriver")  # Replace with your chromedriver path
    driver = webdriver.Chrome(service=service, options=chrome_options)
    return driver

# Function to load HTML and detect alerts
def check_for_alert(html_content):
    driver = setup_driver()
    try:
        # Load the HTML content
        driver.get("data:text/html;charset=utf-8," + html_content)

        # Check for an alert
        try:
            alert = driver.switch_to.alert  # Switch to alert if present
            alert_text = alert.text  # Get the alert's text
            alert.dismiss()  # Close the alert
            #print(f"Alert detected! Content: {alert_text}")
            # We noticed an alert!!!
            # Just do the bullshit.
            fh = open("exploit.txt", "w")
            fh.write(html_content)
            fh.close()
            fh = open("result.txt", "w")
            fh.write("1")
            fh.close()
            return alert_text
        except NoAlertPresentException:
            #print("No alert detected.")
            fh = open("result.txt", "w")
            fh.write("0")
            fh.close()
            return None
    finally:
        driver.quit()  # Clean up and close the browser

# Example HTML/JavaScript with an alert
html_with_alert = """
<!DOCTYPE html>
<html>
<head>
    <title>Test Alert</title>
</head>
<body>
    <script>
        alert("This is a test alert!");
    </script>
</body>
</html>
"""

# Test the function

fh = open("cur_payload.txt", "r")
payload = fh.read()
fh.close()

check_for_alert(payload)
