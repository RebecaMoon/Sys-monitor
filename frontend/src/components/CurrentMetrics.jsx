function CurrentMetrics({ metric }) {
  return (
    <section className="current-metrics">
      <h2>Current Server Usage</h2>

      <div className="metric-cards">
        <div className="metric-card">
          <span className="metric-label">CPU</span>
          <strong className="metric-value">{metric.cpu_percent}%</strong>
        </div>
        <div className="metric-card">
          <span className="metric-label">RAM</span>
          <strong className="metric-value">{metric.memory_percent}%</strong>
        </div>
        <div className="metric-card">
          <span className="metric-label">Disk</span>
          <strong className="metric-value">{metric.disk_percent}%</strong>
        </div>
      </div>
    </section>
  )
}

export default CurrentMetrics