import { useEffect, useState } from 'react'
import './App.css'
import CurrentMetrics from './components/CurrentMetrics.jsx'
import MetricsGraph from './components/MetricsGraph.jsx'
import GraphSection from './components/GraphSection.jsx'

function App() {

  const [latestMetric, setLatestMetric] = useState({
    cpu_percent: 0,
    memory_percent: 0,
    disk_percent: 0,
  })

  async function fetchLatestMetric() {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/metrics/latest`)
    const data = await response.json()

    setLatestMetric(data)
  }


  const [metricsHistory, setMetricsHistory] = useState([])

  async function fetchMetricsHistory() {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/metrics/history`)

      if (!response.ok) {
        throw new Error("Failed to fetch metrics history")
      }

      const data = await response.json()

      setMetricsHistory(data)

    } catch (error) {
      console.error(error)
    }
  }

  useEffect(() => {

    fetchLatestMetric()
    fetchMetricsHistory()

    const interval = setInterval(() => {

      fetchLatestMetric()
      fetchMetricsHistory()

    }, 3000)

    return () => clearInterval(interval)

  }, [])






  
  return (
    <main className="dashboard">
      <section className="metrics-panel">
        <CurrentMetrics metric={latestMetric} />
        <GraphSection history={metricsHistory} />
      </section>
    </main>
  )
}

export default App