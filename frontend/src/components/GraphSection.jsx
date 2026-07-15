import MetricsGraph from "./MetricsGraph"

function GraphSection({history}) {
  return (
    <section className="graph-section">
      <h2>Usage History</h2>

      <div className="graphs-container">
        <MetricsGraph
          title="CPU"
          history={history}
          dataKey="cpu_percent"
        />

        <MetricsGraph
            title="RAM"
            history={history}
            dataKey="memory_percent"
        />

        <MetricsGraph
            title="Disk"
            history={history}
            dataKey="disk_percent"
        />
      </div>
    </section>
  )
}

export default GraphSection