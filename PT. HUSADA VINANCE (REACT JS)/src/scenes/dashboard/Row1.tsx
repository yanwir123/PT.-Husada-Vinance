import React, { useEffect, useState } from "react";
import {
  ResponsiveContainer,
  CartesianGrid,
  AreaChart,
  BarChart,
  LineChart,
  XAxis,
  YAxis,
  Legend,
  Line,
  Tooltip,
  Area,
  Bar,
} from "recharts";
import DashboardBox from "@/components/DashboardBox";
import BoxHeader from "@/components/BoxHeader";
import { useTheme } from "@mui/material";
import "./Row1.css";

interface KeuanganData {
  Bulan: string;
  Hari: string;
  "Transaksi Masuk": number;
  "Transaksi Keluar": number;
  "Jumlah Transaksi": number;
  "Total Transaksi": number;
  Total: number;
}

const Row1 = () => {
  const { palette } = useTheme();
  const [data, setData] = useState<KeuanganData[]>([]);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch(
        "http://localhost:8080/api/PT.HusadaVinance/GetKeuangan"
      );
      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }
      const jsonData = await response.json();
      setData(jsonData.Data);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  const revenueExpensesData = data.map(
    ({
      Hari,
      "Transaksi Masuk": transaksiMasuk,
      "Transaksi Keluar": transaksiKeluar,
    }) => ({
      name: Hari,
      revenue: transaksiMasuk,
      expenses: transaksiKeluar,
    })
  );

  const revenueProfitData = data.map(
    ({
      Bulan,
      "Jumlah Transaksi": jumlahTransaksi,
      "Total Transaksi": totalTransaksi,
    }) => ({
      name: Bulan.substring(0, 3),
      revenue: ((jumlahTransaksi / totalTransaksi) * 100).toFixed(2), // Mengubah menjadi persentase
      profit: (
        ((jumlahTransaksi - totalTransaksi) / totalTransaksi) *
        100
      ).toFixed(2), // Mengubah menjadi persentase
    })
  );

  const revenueData = data.map(({ Bulan, Total: transaksiMasuk }) => ({
    name: Bulan.substring(0, 3),
    Total: transaksiMasuk,
  }));

  return (
    <>
      <DashboardBox gridArea="a">
        <BoxHeader
          title="Data Week"
          subtitle="Dihitung dalam beberapa minggu terakhir"
          sideText="+4%"
        />
        <ResponsiveContainer width="100%" height={400}>
          <AreaChart
            width={500}
            height={400}
            data={revenueExpensesData}
            margin={{ top: 15, right: 25, left: -10, bottom: 60 }}
          >
            <defs>
              <linearGradient id="colorRevenue" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor={palette.primary.main}
                  stopOpacity={0.5}
                />
                <stop
                  offset="95%"
                  stopColor={palette.primary.main}
                  stopOpacity={0}
                />
              </linearGradient>
              <linearGradient id="colorExpenses" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor={palette.secondary.main}
                  stopOpacity={0.5}
                />
                <stop
                  offset="95%"
                  stopColor={palette.secondary.main}
                  stopOpacity={0}
                />
              </linearGradient>
            </defs>
            <CartesianGrid vertical={false} stroke={palette.grey[300]} />
            <XAxis
              dataKey="name"
              tickLine={false}
              style={{ fontSize: "10px" }}
            />
            <YAxis
              tickLine={false}
              axisLine={{ strokeWidth: "0" }}
              style={{ fontSize: "10px" }}
              domain={[0, "auto"]}
            />
            <Tooltip />
            <Area
              type="monotone"
              dataKey="revenue"
              stroke={palette.primary.main}
              fillOpacity={1}
              fill="url(#colorRevenue)"
            />
            <Area
              type="monotone"
              dataKey="expenses"
              stroke={palette.secondary.main}
              fillOpacity={1}
              fill="url(#colorExpenses)"
            />
          </AreaChart>
        </ResponsiveContainer>
      </DashboardBox>

      <DashboardBox gridArea="b">
        <BoxHeader
          title="Data Mount"
          subtitle="Di Dapat Dari Beberapa Bulan Di Tahun 2022"
          sideText="+124%"
        />
        <ResponsiveContainer width="100%" height={400}>
          <LineChart
            width={500}
            height={400}
            data={revenueProfitData}
            margin={{ top: 20, right: 0, left: -10, bottom: 55 }}
          >
            <CartesianGrid vertical={false} stroke={palette.grey[300]} />
            <XAxis
              dataKey="name"
              tickLine={false}
              style={{ fontSize: "10px" }}
            />
            <YAxis
              yAxisId="left"
              tickLine={false}
              axisLine={false}
              style={{ fontSize: "10px" }}
              tickFormatter={(value) => `${value}%`} // Menambahkan % pada sumbu Y
            />
            <YAxis
              yAxisId="right"
              orientation="right"
              tickLine={false}
              axisLine={false}
              style={{ fontSize: "10px" }}
              tickFormatter={(value) => `${value}%`} // Menambahkan % pada sumbu Y
            />
            <Tooltip formatter={(value, name, props) => `${value}%`} />{" "}
            {/* Menambahkan % pada tooltip */}
            <Legend wrapperStyle={{ margin: "0 0 10px 0" }} />
            <Line
              yAxisId="left"
              type="monotone"
              dataKey="profit"
              stroke={palette.primary.main}
            />
            <Line
              yAxisId="right"
              type="monotone"
              dataKey="revenue"
              stroke={palette.secondary.main}
            />
          </LineChart>
        </ResponsiveContainer>
      </DashboardBox>

      <DashboardBox gridArea="c">
        <BoxHeader
          title="Sub Total PT. Husada Vinance"
          subtitle="Ini Adalah Data Semua Transaksi Perusahaan"
          sideText="+14% Dari Tahun Lalu"
        />
        <ResponsiveContainer width="100%" height={300}>
          <BarChart
            width={500}
            height={300}
            data={revenueData}
            margin={{ top: 17, right: 15, left: -5, bottom: 58 }}
          >
            <CartesianGrid vertical={false} stroke={palette.grey[300]} />
            <XAxis
              dataKey="name"
              axisLine={false}
              tickLine={false}
              style={{ fontSize: "10px" }}
            />
            <YAxis
              axisLine={false}
              tickLine={false}
              style={{ fontSize: "10px" }}
            />
            <Tooltip />
            <Bar dataKey="Total" fill={palette.primary.main} />
          </BarChart>
        </ResponsiveContainer>
      </DashboardBox>
    </>
  );
};

export default Row1;
