import psutil
import time
import requests

# Define the future API URL
API_URL = "http://localhost:8000/metrics"

def get_system_metrics():
    #Collects CPU, Memory, and Disk usage.
    
    metrics = {
        # Save the data in a dictionary 
        "cpu_percent": psutil.cpu_percent(interval=1),
        "memory_percent": psutil.virtual_memory().percent,
        "disk_percent": psutil.disk_usage('/').percent,
        "timestamp": time.time()
    }
    return metrics

def run_agent():
    print("Agent started. Collecting metrics from server...")
    try:
        while True:
            data = get_system_metrics()
            # Print to console for now so we can see it working
            print(f"CPU: {data['cpu_percent']}%")
            print(f"Memory: {data['memory_percent']}%")
            print(f"Disk: {data['disk_percent']}")
            print(f"Time: {data['timestamp']}")
            

            # Send the metrics to the API.
            try:
            
                response = requests.post(API_URL, json=data, timeout=5) # With the timeout 5 the agent doesn't hang if the API doesn't answer
                print(f"Sent data! Status Code: {response.status_code}")
                print("------------------------------")
            except requests.exceptions.ConnectionError:
                print("Failed to send the metrics")
                print("------------------------------")
            
            # Wait 5 seconds before next collection
            time.sleep(5)
    except KeyboardInterrupt:
        print("\nAgent stopped.")

if __name__ == "__main__":
    run_agent()