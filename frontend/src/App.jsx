import './App.css'
import CurrentMetrics from './components/CurrentMetrics.jsx'
import MetricsGraph from './components/MetricsGraph.jsx'

function App() {
  
  const latestMetric = {
    cpu_percent: 24.5,
    memory_percent: 58.2,
    disk_percent: 71.4,
  }

  return (
    <main className="dashboard">
      <section className="metrics-panel">
        <CurrentMetrics metric={latestMetric} />
        <MetricsGraph />
      </section>
    </main>
  )
}

export default App